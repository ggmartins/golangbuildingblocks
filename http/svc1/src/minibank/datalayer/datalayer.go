package minibank

import "io"

type DataLayer interface {
	Init()
	PrintAccounts(io.Writer) error
	PrintTransfers(io.Writer) error
}
