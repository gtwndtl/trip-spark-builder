
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
      script.src = 'https://api.longdo.com/map/?key=YOUR_LONGDO_MAP_KEY';
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
    
    // Create map instance
    mapInstanceRef.current = new window.longdo.Map({
      placeholder: mapRef.current,
      zoom: 13,
      lastView: false,
      location: { lon: 100.5018, lat: 13.7563 }, // Bangkok coordinates
    });
    
    // Add a marker at Bangkok
    const marker = new window.longdo.Marker({ 
      lon: 100.5018, 
      lat: 13.7563 
    }, {
      title: 'กรุงเทพมหานคร',
      detail: 'เมืองหลวงของประเทศไทย',
      visibleRange: { min: 7, max: 19 },
    });
    
    mapInstanceRef.current.Overlays.add(marker);
    
    // Set up map controls and layers
    mapInstanceRef.current.Layers.setBase(window.longdo.Layers.NORMAL);
    mapInstanceRef.current.Layers.add(window.longdo.Layers.TRAFFIC);
    mapInstanceRef.current.Ui.Zoombar.visible(true);
    mapInstanceRef.current.Ui.Toolbar.visible(false);
    mapInstanceRef.current.Ui.DPad.visible(true);
    mapInstanceRef.current.Ui.Crosshair.visible(false);
    mapInstanceRef.current.Ui.Scale.visible(true);
  };

  return (
    <div className="p-4">
      <h3 className="text-lg font-medium mb-3">แผนที่</h3>
      <div 
        ref={mapRef} 
        className="h-[250px] rounded-lg shadow-md overflow-hidden"
      ></div>
    </div>
  );
};

export default LongdoMap;
