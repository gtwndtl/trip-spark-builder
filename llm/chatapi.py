from fastapi import FastAPI
from pydantic import BaseModel
import requests
import json
import re
from fastapi.responses import JSONResponse
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()

# ตั้งค่า CORS ถ้าต้องการเรียก API จาก frontend ที่อยู่คนละ origin
origins = [
    "http://localhost:8080",
    "http://127.0.0.1:8080",
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,  # เปลี่ยนตาม URL frontend ของคุณ
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

class TripRequest(BaseModel):
    message: str

@app.post("/plan-trip")
async def plan_trip(req: TripRequest):
    prompt = f"""
สกัดข้อมูลทริปจากข้อความต่อไปนี้ และตอบกลับในรูปแบบ JSON เท่านั้น:

"{req.message}"

รูปแบบ JSON:
{{
  "destination": "...",
  "days": จำนวนวัน,
  "style": "...",
  "budget": จำนวนเงิน
}}
"""
    try:
        ollama_response = requests.post(
            "http://localhost:11434/api/generate",
            json={
                "model": "gemma3:latest",
                "prompt": prompt,
                "stream": False
            }
        )
        ollama_response.raise_for_status()
        content = ollama_response.json().get("response", "")

        # ลบเครื่องหมาย ```json และ ``` ออกจากข้อความก่อนแปลง JSON
        cleaned = re.sub(r"```json|```", "", content).strip()

        # แปลง string JSON ที่ได้เป็น dict
        parsed_json = json.loads(cleaned)

    except requests.RequestException as e:
        return JSONResponse(status_code=500, content={"error": "Failed to call Ollama API", "detail": str(e)})
    except json.JSONDecodeError:
        return JSONResponse(status_code=500, content={"error": "Ollama response is not valid JSON", "response": content})

    return {"parsed": parsed_json}


class Trip(BaseModel):
    destination: str
    duration: str
    budget: str
    style: str

@app.post("/trip-planner")
async def trip_planner(req: Trip):
    prompt = f"""
ช่วยวางแผนการท่องเที่ยวจากข้อมูลต่อไปนี้:
- สถานที่: {req.destination}
- ระยะเวลา: {req.duration} วัน
- งบประมาณ: {req.budget or "ไม่ระบุ"} บาท
- สไตล์การท่องเที่ยว: {req.style or "ทั่วไป"}

แผนการเดินทางควรมีรายละเอียดดังนี้:
1. ตารางกิจกรรมรายวัน พร้อมเวลา
2. ร้านอาหารแนะนำในแต่ละมื้อ
3. ที่พักที่เหมาะกับงบประมาณ
4. วิธีการเดินทางระหว่างสถานที่
5. ประมาณการค่าใช้จ่ายแต่ละวัน

ตอบในรูปแบบ JSON ที่มีโครงสร้างดังนี้:

{{
  "days": [
    {{
      "day": 1,
      "date": "วันที่",
      "activities": [
        {{
          "time": "เวลา",
          "title": "ชื่อกิจกรรม",
          "description": "รายละเอียด"
        }}
      ],
      "routes": [
        {{
          "from": "ต้นทาง",
          "to": "ปลายทาง",
          "transport": "วิธีเดินทาง",
          "duration": "เวลาเดินทาง"
        }}
      ]
    }}
  ],
  "budget_breakdown": {{
    "accommodation": จำนวนเงิน,
    "food": จำนวนเงิน,
    "transport": จำนวนเงิน,
    "activities": จำนวนเงิน,
    "total": จำนวนเงิน
  }}
}}
"""


    try:
        ollama_response = requests.post(
            "http://localhost:11434/api/generate",
            json={
                "model": "gemma3:latest",
                "prompt": prompt,
                "stream": False
            }
        )
        ollama_response.raise_for_status()
        content = ollama_response.json().get("response", "")

        return {"plan": content}  # ✅ แก้ตรงนี้ ไม่ต้องแปลง JSON

    except requests.RequestException as e:
        return JSONResponse(status_code=500, content={"error": "Failed to call Ollama API", "detail": str(e)})