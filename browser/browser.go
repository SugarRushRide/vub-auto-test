// 浏览器实例的生命周期管理（创建/销毁）
// 浏览器配置（无头模式、窗口尺寸、代理等）
// 提供浏览器上下文（context）
// Created: 2025/4/15

package browser

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"time"
	"vub-auto-test/config"
)

// Browser 封装浏览器实例和配置
type Browser struct {
	ctx    context.Context    // 浏览器操作上下文
	cancel context.CancelFunc // 取消函数用于关闭浏览器
	config config.Browser     // 从主配置继承的浏览器配置
}

// NewBrowser 创建浏览器实例
func NewBrowser(browserConfig config.Browser) (*Browser, error) {
	// 1.设置浏览器启动参数
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", browserConfig.Headless),
		chromedp.DisableGPU,
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.UserAgent(browserConfig.UserAgent),
		chromedp.WindowSize(
			browserConfig.WindowSize.Width,
			browserConfig.WindowSize.Height,
		),
	)

	// 2. 指定Chrome可执行路径（如果配置中存在）
	if browserConfig.ExecPath != "" {
		opts = append(opts, chromedp.ExecPath(browserConfig.ExecPath))
	}

	// 3. 配置代理（如果存在）
	if browserConfig.Proxy != "" {
		opts = append(opts, chromedp.ProxyServer(browserConfig.Proxy))
	}

	// 4. 创建浏览器上下文
	allocCtx, cancelAlloc := chromedp.NewExecAllocator(
		context.Background(),
		opts...,
	)

	// 5. 创建浏览器上下文
	ctx, cancelCtx := chromedp.NewContext(allocCtx)

	// 6. 创建带超时的上下文
	ctx, cancelTimeout := context.WithTimeout(
		ctx,
		browserConfig.Timeout,
	)

	// 合并取消函数
	cancel := func() {
		cancelCtx()
		cancelTimeout()
		cancelAlloc()
	}

	// 6. 验证浏览器是否可用
	if err := chromedp.Run(ctx); err != nil {
		cancel()
		return nil, fmt.Errorf("Browser Run: %w", err)
	}

	return &Browser{
		ctx:    ctx,
		cancel: cancel,
		config: browserConfig,
	}, nil
}

// Close 关闭浏览器
func (b *Browser) Close() {
	// 1.取消上下文释放
	b.cancel()

	// 2. 等待资源释放
	time.Sleep(500 * time.Millisecond)
}

// Context 暴露浏览器上下文
func (b *Browser) Context() context.Context {
	return b.ctx
}

// Config 获取浏览器配置
func (b *Browser) Config() config.Browser {
	return b.config
}
