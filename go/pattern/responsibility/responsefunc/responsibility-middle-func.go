package responsefunc

//middle responsefunc

type MiddleFunc func() error

type Pipeline struct {
	methods []MiddleFunc
	finally MiddleFunc
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) ChildPipeline(pipeline Pipeline) *Pipeline {
	// 追加
	p.methods = append(p.methods, func() error {
		return pipeline.Run()
	})
	// 处理
	return p
}

func (p *Pipeline) Add(methods MiddleFunc) *Pipeline {
	p.methods = append(p.methods, methods)
	return p
}

func (p *Pipeline) AddLink(link *Link) *Pipeline {
	// 追加
	p.methods = append(p.methods, func() error {
		return link.Run()
	})
	return p
}



func (p *Pipeline) Run() error {
	// 定义错误
	var err error
	// 遍历循环
	for _, method := range p.methods {
		// 调用执行
		err = method()
	}
	// 最终执行
	if p.finally != nil {
		_ = p.finally()
	}
	// 返回
	return err
}


