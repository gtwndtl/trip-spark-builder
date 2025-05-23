
export const tripParsingPrompt = `
สกัดข้อมูลทริปจากข้อความ และตอบกลับในรูปแบบ JSON เท่านั้น

รูปแบบ JSON ที่ต้องการ:
{
  "destination": "ชื่อสถานที่",
  "days": จำนวนวัน,
  "style": "สไตล์การท่องเที่ยว",
  "budget": จำนวนเงิน
}

แนวทางในการสกัดข้อมูล:
1. สถานที่: ระบุชื่อจังหวัด เมือง หรือประเทศ ที่ชัดเจน
2. จำนวนวัน: ระบุเป็นตัวเลขเท่านั้น (เช่น 2, 3, 5)
3. สไตล์: อธิบายรูปแบบการท่องเที่ยวเป็นวลีสั้นๆ (เช่น "ชิวๆริมทะเล", "ผจญภัย", "วัฒนธรรม")
4. งบประมาณ: ระบุเป็นตัวเลขเท่านั้น ไม่ต้องใส่สัญลักษณ์สกุลเงิน

หากไม่มีข้อมูลใดในข้อความ ให้ใช้ค่า null
`;

export const tripPlanningPrompt = `
ช่วยวางแผนการท่องเที่ยวจากข้อมูลต่อไปนี้:
- สถานที่: {destination}
- ระยะเวลา: {days} วัน
- สไตล์การท่องเที่ยว: {style}
- งบประมาณ: {budget} บาท

แผนการเดินทางควรมีรายละเอียดดังนี้:
1. ตารางกิจกรรมรายวัน พร้อมเวลา
2. ร้านอาหารแนะนำในแต่ละมื้อ
3. ที่พักที่เหมาะกับงบประมาณ
4. วิธีการเดินทางระหว่างสถานที่
5. ประมาณการค่าใช้จ่ายแต่ละวัน

ตอบในรูปแบบ JSON ที่มีโครงสร้างดังนี้:
{
  "days": [
    {
      "day": 1,
      "date": "วันที่",
      "activities": [
        {
          "time": "เวลา",
          "title": "ชื่อกิจกรรม",
          "description": "รายละเอียด"
        }
      ],
      "routes": [
        {
          "from": "ต้นทาง",
          "to": "ปลายทาง",
          "transport": "วิธีเดินทาง",
          "duration": "เวลาเดินทาง"
        }
      ]
    }
  ],
  "budget_breakdown": {
    "accommodation": จำนวนเงิน,
    "food": จำนวนเงิน,
    "transport": จำนวนเงิน,
    "activities": จำนวนเงิน,
    "total": จำนวนเงิน
  }
}
`;
