package herder

import "time"

type Option func(*Herder)

func User(u string) Option {
	return func(h *Herder) {
		h.user = u
	}
}

func Owner(o string) Option {
	return func(h *Herder) {
		h.owner = o
	}
}

func Addr(a string) Option {
	return func(h *Herder) {
		h.addr = a
	}
}

func Delay(d time.Duration) Option {
	return func(h *Herder) {
		h.delay = d
	}
}

func Check(c time.Duration) Option {
	return func(h *Herder) {
		h.check = c
	}
}

func Verbose(v bool) Option {
	return func(h *Herder) {
		h.verbose = v
	}
}
