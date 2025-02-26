package util

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func WriteLog(taskName string) (func(string) (bool, error), func() error, error) {
	folderName := taskName
	// 创建多级目录
	if err := os.MkdirAll(folderName, 0755); err != nil {
		return nil, nil, err
	}
	// 生成安全文件名（替换冒号）
	fileName := fmt.Sprintf("%s_%s.log",
		taskName,
		time.Now().Format("20060102-150405"),
	)
	filePath := filepath.Join(folderName, fileName)
	// 以追加模式打开文件（不存在则创建）
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, nil, err
	}

	// 闭包函数：写入日志
	writeFunc := func(logs string) (bool, error) {
		line := fmt.Sprintf("%s   %s\n",
			time.Now().Format(time.DateTime),
			logs,
		)
		if _, err := file.WriteString(line); err != nil {
			return false, err
		}
		return true, nil
	}

	// 关闭文件函数
	closeFunc := func() error {
		if file != nil {
			err := file.Close()
			file = nil // 避免重复关闭
			return err
		}
		return errors.New("file already closed")
	}

	return writeFunc, closeFunc, nil
}

func WriteSuccess(taskName string) (func(string) (bool, error), func() error, error) {
	folderName := taskName
	// 创建多级目录
	if err := os.MkdirAll(folderName, 0755); err != nil {
		return nil, nil, err
	}
	// 生成安全文件名（替换冒号）
	fileName := fmt.Sprintf("Success_%s_%s.log",
		taskName,
		time.Now().Format("20060102-150405"),
	)
	filePath := filepath.Join(folderName, fileName)
	// 以追加模式打开文件（不存在则创建）
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, nil, err
	}

	// 闭包函数：写入日志
	writeFunc := func(logs string) (bool, error) {
		line := fmt.Sprintf("%s   %s\n",
			time.Now().Format(time.DateTime),
			logs,
		)
		if _, err := file.WriteString(line); err != nil {
			return false, err
		}
		return true, nil
	}

	// 关闭文件函数
	closeFunc := func() error {
		if file != nil {
			err := file.Close()
			file = nil // 避免重复关闭
			return err
		}
		return errors.New("file already closed")
	}

	return writeFunc, closeFunc, nil
}
