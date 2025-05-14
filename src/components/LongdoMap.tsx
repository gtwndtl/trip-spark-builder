
import React, { useEffect, useRef } from 'react';

declare global {
  interface Window {
    longdo: any;
    longdoMapCallback: () => void;
  }
}

const LongdoMap = () => {
  const mapRef = useRef<HTMLDivElement>(null);
  const mapInstanceRef = useRef<any>(null);

  useEffect(() => {
    // Load Longdo Map script if it's not already loaded
    if (!window.longdo) {
      const script = document.createElement('script');
      script.src = 'https://api.longdo.com/map/?key=f278aaef2d456a4e85e80715f7f32ef9';
      script.async = true;
      script.defer = true;
      
      // Define callback function
      window.longdoMapCallback = () => {
        initMap();
      };
      
      script.onload = () => {
        // The script has loaded
        if (window.longdo) {
          initMap();
        }
      };
      
      document.head.appendChild(script);
    } else {
      // Longdo Map is already loaded, initialize map
      initMap();
    }
    
    return () => {
      // Cleanup function
      if (mapInstanceRef.current) {
        // No explicit destroy method needed for Longdo Map
        mapInstanceRef.current = null;
      }
    };
  }, []);
  
  const initMap = () => {
    if (!mapRef.current || !window.longdo) return;
    
    // Create map instance with improved options
    mapInstanceRef.current = new window.longdo.Map({
      placeholder: mapRef.current,
      zoom: 13,
      lastView: false,
      location: { lon: 100.5018, lat: 13.7563 }, // Bangkok coordinates
      ui: {
        zoomBar: true,
        toolbar: true,
        scale: true,
        layerSelector: true
      }
    });
    
    // Add multiple markers for popular Bangkok spots
    const landmarks = [
      {
        name: 'วัดพระแก้ว',
        desc: 'วัดพระศรีรัตนศาสดาราม',
        pos: { lon: 100.4923, lat: 13.7516 }
      },
      {
        name: 'สยามพารากอน',
        desc: 'ศูนย์การค้าชั้นนำของกรุงเทพฯ',
        pos: { lon: 100.5347, lat: 13.7466 }
      },
      {
        name: 'ตลาดนัดจตุจักร',
        desc: 'ตลาดนัดสุดสัปดาห์ที่ใหญ่ที่สุดในประเทศไทย',
        pos: { lon: 100.5508, lat: 13.7999 }
      },
      {
        name: 'เยาวราช',
        desc: 'ไชน่าทาวน์ของกรุงเทพฯ',
        pos: { lon: 100.5102, lat: 13.7393 }
      }
    ];
    
    // Add each landmark as a marker
    landmarks.forEach(landmark => {
      const marker = new window.longdo.Marker(landmark.pos, {
        title: landmark.name,
        detail: landmark.desc,
        visibleRange: { min: 7, max: 19 },
        icon: {
          url: 'https://map.longdo.com/mmmap/images/pin_mark.png',
          offset: { x: 12, y: 45 }
        }
      });
      
      mapInstanceRef.current.Overlays.add(marker);
    });
    
    // Add Bangkok area highlight
    const polygon = new window.longdo.Polygon([
      { lon: 100.4503, lat: 13.7200 },
      { lon: 100.5503, lat: 13.7200 },
      { lon: 100.5503, lat: 13.8200 },
      { lon: 100.4503, lat: 13.8200 }
    ], {
      title: 'พื้นที่กรุงเทพมหานครชั้นใน',
      detail: 'บริเวณศูนย์กลางธุรกิจของกรุงเทพฯ',
      lineWidth: 2,
      lineColor: '#9b87f5',
      fillColor: 'rgba(155, 135, 245, 0.3)'
    });
    mapInstanceRef.current.Overlays.add(polygon);
    
    // Set up map controls and layers with traffic information
    mapInstanceRef.current.Layers.setBase(window.longdo.Layers.NORMAL);
    mapInstanceRef.current.Layers.add(window.longdo.Layers.TRAFFIC);
    
    // Enable interactive UI components
    mapInstanceRef.current.Ui.Zoombar.visible(true);
    mapInstanceRef.current.Ui.Toolbar.visible(true);
    mapInstanceRef.current.Ui.DPad.visible(true);
    mapInstanceRef.current.Ui.Crosshair.visible(false);
    mapInstanceRef.current.Ui.Scale.visible(true);
    
    // Add event listener for click on map
    mapInstanceRef.current.Event.bind('click', function(mouseEvent: any) {
      console.log('Map clicked at:', mouseEvent.location);
    });
  };

  return (
    <div className="p-4">
      <h3 className="text-lg font-prompt font-medium mb-3">แผนที่กรุงเทพมหานคร</h3>
      <div 
        ref={mapRef} 
        className="h-[350px] md:h-[400px] rounded-lg shadow-md overflow-hidden"
      ></div>
      <div className="mt-2 text-xs text-gray-500 font-sarabun text-right">
        ข้อมูลแผนที่ © Longdo Map
      </div>
    </div>
  );
};

export default LongdoMap;
