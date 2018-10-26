package homework

import (
    "math/rand"
)

type Nginx interface {
    DoSend(hs []*Hosts) *Hosts
}

type Hosts struct {
    IP   string
    PORT string
}

type R struct {
}

func (r *R) DoSend(hs []*Hosts) *Hosts {

    // error 没写

    lens := len(hs)
    index := rand.Intn(lens)

    return hs[index]

}

type OBO struct {
    nowIndex int
}

func (o *OBO) DoSend(hs []*Hosts) *Hosts {

    if o.nowIndex >= len(hs) {
        o.nowIndex = 0
    }
    oh := hs[o.nowIndex]
    o.nowIndex += 1

    return oh
}
