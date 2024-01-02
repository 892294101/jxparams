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

func (p *ParamsSet) GetParams(s string) (*Params, bool) {
	if p.paramKeyValue == nil || len(p.paramSort) == 0 {
		return nil, false
	}

	value := strings.ToLower(s)
	v, ok := p.paramKeyValue[value]
	if ok {
		return v, true
	}
	return nil, false
}

func (p *ParamsSet) SetParams(s string, c ...*Config) {
	value := strings.ToLower(s)
	if p.paramKeyValue == nil {
		p.paramKeyValue = map[string]*Params{}
	}
	if !p.check(value) {
		if len(c) > 0 {
			p.paramKeyValue[value] = &Params{config: c[0]}
		} else {
			p.paramKeyValue[value] = &Params{}
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

	for _, s := range ps.Keys() {
		if r, ok := ps.Get(s); ok {
			pSet[strings.ToLower(s)] = r
		} else {
			pSet[strings.ToLower(s)] = ""
		}
	}

	for key, _ := range pSet {
		_, ok := p.paramKeyValue[strings.ToLower(key)]
		if !ok {
			return fmt.Errorf("%s non prefabricated parameters are illegal", key)
		}
	}

	for key, strc := range p.paramKeyValue {
		if strc.config != nil && strc.config.getMust() {
			v, ok := pSet[key]
			if ok {
				p.paramKeyValue[key].value = strings.TrimSpace(v)
			} else {
				return fmt.Errorf("%s parameter must be set", key)
			}
		}
		if strc.config != nil && len(strc.config.getDefault()) > 0 {
			v, ok := pSet[key]
			if ok {
				p.paramKeyValue[key].value = strings.TrimSpace(v)
			} else {
				p.paramKeyValue[key].value = strings.TrimSpace(strc.config.getDefault())
			}
		}
		if strc.config == nil {
			v, ok := pSet[key]
			if ok {
				p.paramKeyValue[key].value = strings.TrimSpace(v)
			} else {
				p.paramKeyValue[key].value = ""
			}
		}

	}

	return nil

}

func NewParams() *ParamsSet {
	return new(ParamsSet)
}
