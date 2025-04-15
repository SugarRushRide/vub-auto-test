// 页面操作封装：browser/actions.go
// 功能：封装浏览器的常用页面行为，例如打开页面、点击按钮、等待元素等
// Created: 2025/4/15

package browser

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"time"
)

// OpenPage 打开指定URL页面
func OpenPage(ctx context.Context, url string) error {
	fmt.Println("Opening page", url)
	return chromedp.Run(ctx,
		chromedp.Navigate(url),
	)
}

// Click 点击指定元素(通过CSS Selector或XPath)
func Click(ctx context.Context, selector string) error {
	fmt.Println("Clicking: ", selector)
	return chromedp.Run(ctx,
		chromedp.WaitVisible(selector, chromedp.BySearch),
		chromedp.Click(selector, chromedp.BySearch),
	)
}

// WaitVisible 等待元素可见(可用于判断页面是否加载完成)
func WaitVisible(ctx context.Context, selector string, timeout time.Duration) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	fmt.Println("Waiting for ", selector)
	return chromedp.Run(ctxWithTimeout,
		chromedp.WaitVisible(selector, chromedp.BySearch),
	)
}

// InputText 向输入框中输入文字
func InputText(ctx context.Context, selector, text string) error {
	fmt.Println("InputText: ", selector)
	return chromedp.Run(ctx,
		chromedp.WaitVisible(selector, chromedp.BySearch),
		chromedp.SetValue(selector, text, chromedp.BySearch),
	)
}

// GetText 获取元素的文本内容
func GetText(ctx context.Context, selector string) (string, error) {
	var text string
	err := chromedp.Run(ctx,
		chromedp.Text(selector, &text, chromedp.BySearch),
	)
	if err != nil {
		return "", err
	}
	return text, nil
}
