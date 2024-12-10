package jxparams

import (
	"fmt"
	"github.com/magiconair/properties"
	"strconv"
	"strings"
)

type Params struct {
	value  string
	config *Config
}

func (p *Params) ToString() string {
	return p.value
}

func (p *Params) Split(sep string) []string {
	return strings.Split(p.value, sep)
}

func (p *Params) ToInt() (int, error) {
	if len(p.value) == 0 {
		return 0, nil
	}

	num, err := strconv.Atoi(p.value)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (p *Params) ToInt64() (int64, error) {
	if len(p.value) == 0 {
		return 0, nil
	}

	num, err := strconv.ParseInt(p.value, 10, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (p *Params) ToBool() (bool, error) {
	b, err := strconv.ParseBool(p.value)
	if err != nil {
		return false, err
	}
	return b, nil
}

func (p *Params) ToFloat64() (float64, error) {
	if len(p.value) == 0 {
		return 0, nil
	}

	num, err := strconv.ParseFloat(p.value, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

type ParamsSet struct {
	filePathSrc   string
	charset       properties.Encoding
	paramSort     []string
	paramKeyValue map[string]*Params
}

func (p *ParamsSet) SetCharset(c properties.Encoding) {
	p.charset = c
}

func (p *ParamsSet) SetConfigFile(s string) {
	p.filePathSrc = s
}

func (p *ParamsSet) check(s string) bool {
	_, ok := p.paramKeyValue[s]
	if ok {
		return true
	}
	return false
}

// GetPrefix 获取前缀参数
func (p *ParamsSet) GetPrefix(s string) (r []*Params, ok bool) {
	if !strings.HasPrefix(s, ".") {
		s = strings.ToLower(s) + "."
	}

	if p.paramKeyValue == nil || len(p.paramKeyValue) == 0 {
		return nil, false
	}

	for s2, params := range p.paramKeyValue {
		if strings.HasPrefix(s2, s) {
			r = append(r, params)
		}
	}
	if len(r) > 0 {
		return r, true
	}
	return nil, false
}

// GetSuffix 获取后缀参数
func (p *ParamsSet) GetSuffix(s string) (r []*Params, ok bool) {
	if !strings.HasSuffix(s, ".") {
		s = "." + strings.ToLower(s)
	}

	if p.paramKeyValue == nil || len(p.paramKeyValue) == 0 {
		return nil, false
	}

	for s2, params := range p.paramKeyValue {
		if strings.HasSuffix(s2, s) {
			r = append(r, params)
		}
	}
	if len(r) > 0 {
		return r, true
	}
	return nil, false
}

func (p *ParamsSet) GetParams(s string) (*Params, bool) {
	if p.paramKeyValue == nil || len(p.paramKeyValue) == 0 {
		return nil, false
	}

	value := strings.ToLower(s)
	v, ok := p.paramKeyValue[value]
	if ok {
		if len(v.value) > 0 {
			return v, true
		} else {
			return nil, false
		}
	}
	return nil, false
}

func (p *ParamsSet) SetParams(s string, c ...*Config) {
	for _, config := range c {
		if config.prefix {
			if !strings.HasPrefix(s, ".") {
				s = s + "."
			}
		}

		if config.suffix {
			if !strings.HasSuffix(s, ".") {
				s = "." + s
			}
		}

	}

	value := strings.ToLower(s)
	if p.paramKeyValue == nil {
		p.paramKeyValue = map[string]*Params{}
	}
	if !p.check(value) {
		if len(c) > 0 {
			p.paramKeyValue[value] = &Params{config: c[0]}
		} else {
			p.paramKeyValue[value] = &Params{config: &Config{}}
		}
		p.paramSort = append(p.paramSort, value)
	}
}

func (p *ParamsSet) Println() {
	if p.paramKeyValue != nil && len(p.paramKeyValue) > 0 && len(p.paramSort) > 0 {
		for _, key := range p.paramSort {
			v, ok := p.paramKeyValue[key]
			if ok {
				fmt.Printf("%v=%v\n", key, v.value)
			}
		}
	}
}

func (p *ParamsSet) Load() error {
	p.charset = properties.UTF8
	if len(p.filePathSrc) == 0 {
		return fmt.Errorf("parameter file path must be set")
	}

	if p.paramKeyValue == nil || len(p.paramKeyValue) == 0 {
		return fmt.Errorf("parameters must be defined")
	}

	ps, err := properties.LoadFile(p.filePathSrc, p.charset)
	if err != nil {
		return err
	}

	pSet := map[string]string{}
	//	var wSet []*Params

	// 把文件中的参数加载到内存中
	for _, s := range ps.Keys() {
		if r, ok := ps.Get(s); ok {
			pSet[strings.ToLower(s)] = r
		} else {
			pSet[strings.ToLower(s)] = ""
		}
	}

	// 匹配参数是否在定义的参数中.
	clearSet := map[string]string{}
	for key, _ := range pSet {
		_, ok := p.paramKeyValue[strings.ToLower(key)] // 精确匹配
		if !ok {
			// 如果精确的匹配不到. 则匹配统配的参数
			var e bool
			for _, v := range p.paramSort {
				if (p.paramKeyValue[v].config.prefix || p.paramKeyValue[v].config.suffix) && (strings.HasPrefix(key, v) || strings.HasSuffix(key, v)) {
					p.paramKeyValue[key] = &Params{value: p.paramKeyValue[v].value, config: p.paramKeyValue[v].config} // 把统配到的参数赋值给新的参数.
					p.paramKeyValue[key].value = pSet[key]                                                             // 把就参数值, 赋值给新参数
					//p.paramSort[i] = key
					//delete(p.paramKeyValue, v).
					p.paramSort = append(p.paramSort, key) // 然后把原来旧的参数删除. 只保留新参数
					clearSet[v] = ""
					e = true
					break
				}
				/*if !p.paramKeyValue[v].config.prefix && p.paramKeyValue[v].config.suffix && strings.HasSuffix(key, v) {
					p.paramKeyValue[key] = p.paramKeyValue[v]
					p.paramKeyValue[key].value = pSet[key]
					p.paramSort[i] = key
					delete(p.paramKeyValue, v)
					e = true
					break
				}*/
			}
			if !e {
				return fmt.Errorf("%s non prefabricated parameters are illegal", key)
			}
		}
	}

	for key, _ := range clearSet {
		delete(p.paramKeyValue, key)
		for i, s := range p.paramSort {
			if strings.EqualFold(s, key) {
				p.paramSort = append(p.paramSort[:i], p.paramSort[i+1:]...)
				break
			}
		}
	}

	// 效验参数. 按照配置.
	// 必须参数必须存在.
	// 如果设置默认值, 那么当参数不存在时, 将会自动设置参数
	for key, strc := range p.paramKeyValue {
		// 如果参数的配置为nil, 则把文件中的参数值直接赋值.
		if strc.config == nil {
			v, ok := pSet[key]
			if ok {
				p.paramKeyValue[key].value = strings.TrimSpace(v)
			} else {
				p.paramKeyValue[key].value = ""
			}
		}

		// 判断参数是否为必须存在
		if strc.config != nil && strc.config.getMust() {
			v, ok := pSet[key]
			if ok {
				p.paramKeyValue[key].value = strings.TrimSpace(v)
			} else {
				return fmt.Errorf("%s parameter must be set", key)
			}
		}
		// 非必须
		if strc.config != nil && !strc.config.getMust() {
			v, ok := pSet[key]
			if ok {
				// 查看是否存在值, 如果参数文件存在值, 则直接赋值
				if len(strings.TrimSpace(v)) > 0 {
					p.paramKeyValue[key].value = strings.TrimSpace(v)
				} else {
					// 如果参数文件不存在值, 则查看参数是否设置了默认值. 如果存在默认值, 则直接赋值.
					if len(strc.config.getDefault()) > 0 {
						p.paramKeyValue[key].value = strings.TrimSpace(strc.config.getDefault())
					}
				}

			}
		}

		/*// 判断参数是否需要设置默认值
		if strc.config != nil && len(strc.config.getDefault()) > 0 {
			v, ok := pSet[key]
			if ok {
				p.paramKeyValue[key].value = strings.TrimSpace(v)
			} else {
				p.paramKeyValue[key].value = strings.TrimSpace(strc.config.getDefault())
			}
		}*/

	}

	return nil

}

func NewParams() *ParamsSet {
	return new(ParamsSet)
}
