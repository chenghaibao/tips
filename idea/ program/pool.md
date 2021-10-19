
# 池化技术

## 池化策略 解决哪些问题

1. 重复使用 
2. 减少创建的代价

## 示例代码

```
// 连接池结构体
type Pool struct {
	ch      chan Conn
	max     int
	num     int
	closed  bool
	mu      sync.Mutex
}

// 新建连接池
func newPool(config *structs.SourceConfig) *Pool {
	return &Pool{
		ch:     make(chan Conn, config.MaxConnNums),
		max:    config.MaxConnNums,
	}
}

// 创建mysql连接
func (p *Pool) newConn() Conn {
	var conn Conn
	switch p.Config.Driver {
	case "mysql":
		conn := &mysql.Conn{
			Config: p.Config,
		}
		conn.Connect()
		// 判断是否存在优先执行的内容
		if p.firstDo != "" {
			// 连接需先执行
			conn.Exec(p.firstDo)
		}
		return conn
	default:
		panic("不支持该驱动")
	}
	return conn
}


// 获取mysql连接
func (p *Pool) Get() Conn {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.ch) == 0 && p.num < p.max {
		p.num++
		return p.newConn()
	} else {
		return <-p.ch
	}
}

// Put 返还连接给连接池
func (p *Pool) Put(conn Conn) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// 连接池满了就关闭当前连接
	if len(p.ch) < p.max {
		p.ch <- conn
	} else {
		_ = conn.Close()
	}
}

// 关闭连接池
func (p *Pool) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed {
		return nil
	}
	p.closed = true
	if p.ch != nil {
		close(p.ch)
	}
	// 关闭连接池所有连接
	for c := range p.ch {
		_ = c.Close()
	}
	return nil
}
```

# 示例维护多个连接池
```
// 连接池管理
type PoolManager struct {
	Config *structs.Config
	pools  map[string]*Pool
}

// 维护多个连接池
func (pm *PoolManager) Get(name string) *Pool {
	// 判断内部有没有
	if _, exist := pm.pools[name]; !exist {
		config := pm.getConfigByName(name)
		pm.pools[name] = newPool(config)
	}
	// 返回
	return pm.pools[name]
}
```
