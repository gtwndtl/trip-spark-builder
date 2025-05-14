
import React, { useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import { useToast } from '@/hooks/use-toast';
import ItinerarySummary from '@/components/ItinerarySummary';
import { FullItinerary } from '@/types';
import Header from '@/components/Header';
import Footer from '@/components/Footer';

const TripSummary = () => {
  const location = useLocation();
  const { toast } = useToast();
  const itinerary = location.state?.itinerary as FullItinerary;

  useEffect(() => {
    if (!itinerary) {
      toast({
        title: "ไม่พบข้อมูลแผนการเดินทาง",
        description: "กรุณากลับไปสร้างแผนการเดินทางก่อน",
        variant: "destructive",
      });
    }
  }, [itinerary, toast]);

  const sampleItinerary: FullItinerary = {
    preferences: {
      destination: "กรุงเทพฯ",
      duration: 3,
      budget: "ปานกลาง",
      style: "วัฒนธรรม"
    },
    days: [
      {
        day: 1,
        date: "15 พฤษภาคม 2568",
        activities: [
          {
            time: "09:00",
            title: "เช็คอินที่โรงแรม Grande Centre Point สุขุมวิท 55",
            description: "เริ่มต้นทริปของคุณด้วยการเช็คอินที่โรงแรมย่านสุขุมวิท ใกล้ BTS ทองหล่อ"
          },
          {
            time: "11:30",
            title: "ทานอาหารกลางวันที่ร้าน สมบูรณ์โภชนา",
            description: "ลิ้มลองอาหารไทยต้นตำรับที่เปิดมายาวนานกว่า 50 ปี"
          },
          {
            time: "14:00",
            title: "เที่ยวชมวัดพระแก้ว",
            description: "สัมผัสความงดงามของสถาปัตยกรรมไทยและพระแก้วมรกตที่เป็นสิ่งศักดิ์สิทธิ์คู่บ้านคู่เมือง"
          },
          {
            time: "17:00",
            title: "ล่องเรือชมวิวแม่น้ำเจ้าพระยา",
            description: "ชมวิวทิวทัศน์สองฝั่งแม่น้ำเจ้าพระยา พร้อมรับประทานอาหารบนเรือ"
          },
          {
            time: "20:00",
            title: "เดินเที่ยวที่เยาวราช",
            description: "สัมผัสบรรยากาศไชน่าทาวน์ของกรุงเทพฯ และลิ้มลองอาหารจีนแท้ๆ"
          }
        ],
        routes: [
          { from: "โรงแรม", to: "ร้านสมบูรณ์โภชนา", transport: "แท็กซี่", duration: "25 นาที" },
          { from: "ร้านสมบูรณ์โภชนา", to: "วัดพระแก้ว", transport: "แท็กซี่", duration: "30 นาที" },
          { from: "วัดพระแก้ว", to: "ท่าเรือ", transport: "เดิน", duration: "15 นาที" },
          { from: "ท่าเรือ", to: "เยาวราช", transport: "รถไฟฟ้า MRT", duration: "25 นาที" }
        ]
      },
      {
        day: 2,
        date: "16 พฤษภาคม 2568",
        activities: [
          {
            time: "08:00",
            title: "อาหารเช้าที่โรงแรม",
            description: "เริ่มต้นวันด้วยบุฟเฟ่ต์อาหารเช้านานาชาติที่โรงแรม"
          },
          {
            time: "10:00",
            title: "ตลาดน้ำดำเนินสะดวก",
            description: "สัมผัสประสบการณ์ตลาดน้ำขึ้นชื่อของไทย"
          },
          {
            time: "14:00",
            title: "อุทยานประวัติศาสตร์พระนครศรีอยุธยา",
            description: "เยี่ยมชมอดีตเมืองหลวงเก่าของไทยที่เป็นมรดกโลก"
          },
          {
            time: "18:30",
            title: "กลับกรุงเทพ - ทานอาหารเย็นที่ Terminal 21",
            description: "ช้อปปิ้งและทานอาหารที่ห้างสรรพสินค้าธีมท่องเที่ยวรอบโลก"
          }
        ],
        routes: [
          { from: "โรงแรม", to: "ตลาดน้ำดำเนินสะดวก", transport: "รถตู้", duration: "1 ชั่วโมง 30 นาที" },
          { from: "ตลาดน้ำดำเนินสะดวก", to: "อยุธยา", transport: "รถตู้", duration: "1 ชั่วโมง 45 นาที" },
          { from: "อยุธยา", to: "Terminal 21", transport: "รถตู้", duration: "1 ชั่วโมง 15 นาที" },
          { from: "Terminal 21", to: "โรงแรม", transport: "BTS", duration: "10 นาที" }
        ]
      },
      {
        day: 3,
        date: "17 พฤษภาคม 2568",
        activities: [
          {
            time: "09:00",
            title: "อาหารเช้าที่โรงแรม",
            description: "เริ่มต้นวันด้วยบุฟเฟ่ต์อาหารเช้านานาชาติที่โรงแรม"
          },
          {
            time: "11:00",
            title: "ช้อปปิ้งที่สยามพารากอน",
            description: "เพลิดเพลินกับการช้อปปิ้งสินค้าแบรนด์เนมและสินค้าไทย"
          },
          {
            time: "14:30",
            title: "พิพิธภัณฑ์ศิลปะไทยร่วมสมัย (MOCA)",
            description: "ชมงานศิลปะร่วมสมัยที่ใหญ่ที่สุดในประเทศไทย"
          },
          {
            time: "18:00",
            title: "อาหารเย็นที่ร้าน เจ๊ไฝ ทองหล่อ",
            description: "ปิดท้ายทริปด้วยอาหารไทย-จีนรสเด็ด"
          },
          {
            time: "21:00",
            title: "เครื่องดื่มที่ Octave Rooftop Bar",
            description: "ชมวิวกรุงเทพฯ ยามค่ำคืนที่บาร์บนดาดฟ้า"
          }
        ],
        routes: [
          { from: "โรงแรม", to: "สยามพารากอน", transport: "BTS", duration: "20 นาที" },
          { from: "สยามพารากอน", to: "MOCA", transport: "แท็กซี่", duration: "25 นาที" },
          { from: "MOCA", to: "ร้านเจ๊ไฝ", transport: "แท็กซี่", duration: "30 นาที" },
          { from: "ร้านเจ๊ไฝ", to: "Octave Rooftop Bar", transport: "เดิน", duration: "15 นาที" },
          { from: "Octave Rooftop Bar", to: "โรงแรม", transport: "เดิน", duration: "10 นาที" }
        ]
      }
    ]
  };

  return (
    <div className="min-h-screen bg-gray-50 flex flex-col">
      <Header />
      
      <main className="flex-1 py-16">
        <div className="container mx-auto px-4">
          <h2 className="text-3xl font-light text-center mb-3">
            <span className="font-normal">สรุปแผนการเดินทาง</span>
          </h2>
          <p className="text-center text-gray-600 mb-8">รายละเอียดการเดินทางและกิจกรรมทั้งหมด</p>
          
          <ItinerarySummary itinerary={itinerary || sampleItinerary} />
          
          <div className="mt-8 flex justify-center">
            <button 
              className="apple-button px-8 py-3 text-white rounded-md hover:bg-opacity-90"
              onClick={() => window.print()}
            >
              พิมพ์แผนการเดินทาง
            </button>
          </div>
        </div>
      </main>
      
      <Footer />
    </div>
  );
};

export default TripSummary;
