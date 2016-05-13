package models

//go:generate ffjson $GOFILE
type Product struct {
        Code        string
        Name        string
        Description string
        Price       float64
        Expiry      int32
        IsPlan      bool
        IsUnlimited bool
        SizeMb      int32
        Is4g        bool
        AutoRenew   bool
        TermsUrl    string
}