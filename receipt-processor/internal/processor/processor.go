package processor

import (
    "errors"
    "fmt"
    "sync"
    "math"
    "strings"
    "time"
    "strconv"
)

// Receipt represents the receipt structure.
type Receipt struct {
    Retailer      string `json:"retailer"`
    PurchaseDate  string `json:"purchaseDate"`
    PurchaseTime  string `json:"purchaseTime"`
    Items         []Item `json:"items"`
    Total         string `json:"total"`
}

// Item represents an item on the receipt.
type Item struct {
    ShortDescription string  `json:"shortDescription"`
    Price            string `json:"price"`
}

var (
    receipts = make(map[string]int)
    mu       sync.Mutex
)

// Process calculates the points for a receipt and stores it.
func Process(receipt *Receipt) (string, error) {
    totalFloat, err := strconv.ParseFloat(receipt.Total, 64)
    if err != nil {
        return "", fmt.Errorf("error converting total to float: %v", err)
    }

    // Calculate points (implementation of calculatePoints function needs total as float64)
    points := calculatePoints(receipt, totalFloat)
    // For demonstration, generate a simple ID and return it.
    id := fmt.Sprintf("%d", time.Now().UnixNano())
    mu.Lock()
    receipts[id] = points 
    mu.Unlock()
    // fmt.println(id,points)
    return id, nil
}

// GetPoints returns the points for a given receipt ID.
func GetPoints(id string) (int, error) {
    mu.Lock()
    points, ok := receipts[id]
    mu.Unlock()
    if !ok {
        // fmt.println(id)
        return 0, errors.New("receipt not found")
    }
    return points, nil
}

// calculatePoints applies the rules to calculate points for a receipt.
func calculatePoints(receipt *Receipt, totalFloat float64) int {
    var totalPoints int = 0

    // One point for every alphanumeric character in the retailer name.
    for _, r := range receipt.Retailer {
        if r >= '0' && r <= '9' || r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z' {
            totalPoints++
        }
    }
    // fmt.println(totalPoints)
    // 50 points if the total is a round dollar amount with no cents.
    if totalFloat == float64(int(totalFloat)) {
        totalPoints += 50
    }
    // fmt.println(totalPoints)
    // 25 points if the total is a multiple of 0.25.
    if int(totalFloat*100)%25 == 0 {
        totalPoints += 25
    }
    // fmt.println(totalPoints)
    // 5 points for every two items on the receipt.
    totalPoints += int((len(receipt.Items) / 2) * 5)
    // fmt.println(totalPoints)
    // Points for item descriptions and prices.
    for _, item := range receipt.Items {
        priceFloat, err := strconv.ParseFloat(item.Price, 64)
        if err != nil {
            // Handle error, perhaps log it, and consider how you want to treat this case.
            // fmt.println("Error converting item price to float:", err)
            continue // Skip this item or return, depending on your error handling strategy.
        }
        trimmedDescription := strings.TrimSpace(item.ShortDescription)
        if len(trimmedDescription)%3 == 0 {
            // Multiply the price by 0.2 and round up to the nearest integer.
            pointsFromPrice := int(math.Ceil(priceFloat * 0.2))
            totalPoints += pointsFromPrice
        }
    }
    // fmt.println(totalPoints)
    
    // 6 points if the day in the purchase date is odd.
    date, err := time.Parse("2006-01-02", receipt.PurchaseDate)
    // fmt.println(date)
    if err == nil && date.Day()%2 != 0 {
        totalPoints += 6
    }
    // fmt.println(totalPoints)
    // 10 points if the time of purchase is after 2:00pm and before 4:00pm.
    timeOfPurchase, err := time.Parse("15:04", receipt.PurchaseTime)
    // fmt.println(timeOfPurchase)
    if err == nil {
        
        purchaseHour := timeOfPurchase.Hour()
        purchaseMinute := timeOfPurchase.Minute()

        // Define start and end times (14:00 to 16:00) in terms of hours and minutes
        startHour, startMinute := 14, 0
        endHour, endMinute := 16, 0

        // Check if time of purchase is within the range
        if (purchaseHour > startHour || (purchaseHour == startHour && purchaseMinute >= startMinute)) &&
        (purchaseHour < endHour || (purchaseHour == endHour && purchaseMinute < endMinute)) {
        // Purchase time is between 2:00 PM and 4:00 PM
        totalPoints += 10
        }
    }
    // fmt.println(totalPoints)
    return totalPoints
    
}

