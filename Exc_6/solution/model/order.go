package model

import (
    "fmt"
    "time"
)

const (
    orderFilename = "order_%d.md"

    // Markdown template for receipts.
    // Fields are populated with fmt.Sprintf:
    //   %d -> Order ID
    //   %s -> CreatedAt formatted timestamp
    //   %d -> DrinkID
    //   %d -> Amount
    markdownTemplate = `# Order: %d

| Created At      | Drink ID | Amount |
|-----------------|----------|--------|
| %s | %d        | %d      |

Thanks for drinking with us!
`
)

type Order struct {
    Base
    Amount  uint64 `json:"amount"`
    DrinkID uint   `json:"drink_id" gorm:"not null"`
    Drink   Drink  `json:"drink" gorm:"foreignKey:DrinkID;references:ID"`
}


// ToMarkdown generates a markdown receipt string for the order.
func (o *Order) ToMarkdown() string {
    return fmt.Sprintf(
        markdownTemplate,
        o.ID,
        o.CreatedAt.Format(time.Stamp),
        o.DrinkID,
        o.Amount,
    )
}

// GetFilename returns the filename for the receipt markdown file.
func (o *Order) GetFilename() string {
    return fmt.Sprintf(orderFilename, o.ID)
}
