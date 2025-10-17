package main

import (
	"fmt"
	"sync"
)

// 调查房屋的信息
type HouseInfo struct {
	ID           int      // 唯一标识
	Owners       []string // 住户
	Price        int      // 价格
	BankAccounts []int    // 房贷情况
	GetMoney     int      // 收到的钱，可能是负值
}

// 调查房子价值
// 1.找物业，插住户是多少人 Owners
// 2.评价房产淘宝法拍 Price
// 3.去银行调查房贷情况，有未还的返回卡号数组 BankAccounts
// 4.计算住户可以拿到的钱 Price-BankMoney
// 5.存入数据库

// 1，2，3可以并发执行，4必须是2，3全部结束才能执行，5必须是4结束后执行

type Response struct {
	data map[string]any
	err  error
}

func sellHouseInfo(id int) (*HouseInfo, error) {
	respch := make(chan Response, 3)
	wg := &sync.WaitGroup{}
	// 1.找物业，插住户是多少人 Owners
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 模拟查询
		owners := []string{"张三", "李四"}
		respch <- Response{data: map[string]any{"owners": owners}, err: nil}
	}()
	// 2.评价房产淘宝法拍 Price
	wg.Add(1)
	go func() {
		defer wg.Done()
		price := 1000000
		respch <- Response{data: map[string]any{"price": price}, err: nil}
	}()
	// 3.去银行调查房贷情况，有未还的返回卡号数组 BankAccounts
	wg.Add(1)
	go func() {
		defer wg.Done()
		bankAccounts := []int{123456, 654321}
		respch <- Response{data: map[string]any{"bankAccounts": bankAccounts}, err: nil}
	}()
	// 等待所有查询完成
	// go func() {
	wg.Wait()
	close(respch)
	// }()
	houseInfo := &HouseInfo{ID: id}
	for resp := range respch {
		if resp.err != nil {
			return nil, resp.err
		}
		for k, v := range resp.data {
			switch k {
			case "owners":
				houseInfo.Owners = v.([]string)
			case "price":
				houseInfo.Price = v.(int)
			case "bankAccounts":
				houseInfo.BankAccounts = v.([]int)
			}
		}
	}
	// 4.计算住户可以拿到的钱 Price-BankMoney
	bankMoney := 0
	for _, acc := range houseInfo.BankAccounts {
		bankMoney += acc // 模拟银行欠款总额
	}
	houseInfo.GetMoney = houseInfo.Price - bankMoney
	// 5.存入数据库
	// 模拟存入数据库
	return houseInfo, nil
}

func main() {
	houseInfo, err := sellHouseInfo(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(houseInfo)
}
