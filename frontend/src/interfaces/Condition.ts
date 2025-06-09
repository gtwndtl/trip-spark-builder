export interface ConditionInterface {
    ID?: number;
    Day?: string; // Use string to represent date in 'YYYY-MM-DD' format // Use string to represent time in 'HH:mm:ss' format
    Price?: number; // Use number for price
    Accommodation?: string; // Accommodation name or ID
    Landmark?: string; // Landmark name or ID
    Style?: string; // Style of the condition (e.g., budget, luxury)
    User_id?: number; // User ID who created the condition
}
// Day    time.Time   
// 	Price  float32
// 	Accommodation string
// 	Landmark string
// 	User_id  uint      