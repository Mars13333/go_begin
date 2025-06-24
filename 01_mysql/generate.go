package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func main() {
	// 获取当前目录
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取当前目录失败: %v\n", err)
		return
	}

	// 扫描10_*到39_*的文件
	var files []string
	for i := 10; i <= 39; i++ {
		pattern := fmt.Sprintf("%02d_*", i) // %02d 将整数 i 格式化为至少2位的十进制数字，不足2位，在前面用 0 填充
		matches, err := filepath.Glob(filepath.Join(currentDir, pattern))
		if err != nil {
			fmt.Printf("扫描文件失败: %v\n", err)
			continue
		}
		files = append(files, matches...)
	}

	// 过滤出.md文件并排序
	var mdFiles []string
	for _, file := range files {
		if strings.HasSuffix(file, ".md") {
			mdFiles = append(mdFiles, filepath.Base(file))
		}
	}
	sort.Strings(mdFiles)

	// 生成目录
	fmt.Println("# MySQL 学习目录")
	fmt.Println()

	for _, file := range mdFiles {
		// 移除.md扩展名
		name := strings.TrimSuffix(file, ".md")

		// 解析文件名获取编号和标题
		parts := strings.SplitN(name, "_", 2)
		if len(parts) != 2 {
			continue
		}

		number := parts[0]
		title := parts[1]

		// 根据编号确定标题级别
		var level string
		switch {
		case strings.HasPrefix(number, "10") || strings.HasPrefix(number, "11") ||
			strings.HasPrefix(number, "12") || strings.HasPrefix(number, "13") ||
			strings.HasPrefix(number, "14") || strings.HasPrefix(number, "15"):
			level = "##" // 开头无空格
		case strings.HasPrefix(number, "16") || strings.HasPrefix(number, "17") ||
			strings.HasPrefix(number, "18") || strings.HasPrefix(number, "19") ||
			strings.HasPrefix(number, "20") || strings.HasPrefix(number, "21") ||
			strings.HasPrefix(number, "22") || strings.HasPrefix(number, "23") ||
			strings.HasPrefix(number, "24") || strings.HasPrefix(number, "25"):
			level = "###" // 开头1个空格
		case strings.HasPrefix(number, "26") || strings.HasPrefix(number, "27") ||
			strings.HasPrefix(number, "28") || strings.HasPrefix(number, "29") ||
			strings.HasPrefix(number, "30") || strings.HasPrefix(number, "31") ||
			strings.HasPrefix(number, "32") || strings.HasPrefix(number, "33") ||
			strings.HasPrefix(number, "34") || strings.HasPrefix(number, "35") ||
			strings.HasPrefix(number, "36") || strings.HasPrefix(number, "37") ||
			strings.HasPrefix(number, "38") || strings.HasPrefix(number, "39"):
			level = "####" // 开头2个空格
		default:
			level = "##"
		}

		// 生成锚点链接
		anchor := generateAnchor(title)
		fmt.Printf("%s [%s](%s#%s)\n", level, title, file, anchor)
		fmt.Println()
	}
}

// 生成锚点链接
func generateAnchor(title string) string {
	// 移除特殊字符，转换为小写，用连字符替换空格
	re := regexp.MustCompile(`[^\w\s-]`)
	anchor := re.ReplaceAllString(title, "")
	anchor = strings.ToLower(anchor)
	anchor = strings.ReplaceAll(anchor, " ", "-")
	return anchor
}
