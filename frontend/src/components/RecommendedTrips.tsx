
import React from 'react';
import { Carousel, CarouselContent, CarouselItem } from '@/components/ui/carousel';
import { AspectRatio } from '@/components/ui/aspect-ratio';
import { Map } from 'lucide-react';
import { useToast } from '@/hooks/use-toast';

type TripCardProps = {
  title: string;
  location: string;
  imageUrl: string;
  days: number;
};

const TripCard = ({ title, location, imageUrl, days }: TripCardProps) => {
  const { toast } = useToast();

  const handleClick = () => {
    toast({
      title: "ทริปถูกเพิ่มในแผนการเดินทาง",
      description: `${title} ได้ถูกเพิ่มในแผนการเดินทางของคุณแล้ว`,
    });
  };

  return (
    <div className="glass-card rounded-xl overflow-hidden shadow-md hover:shadow-lg transition-all cursor-pointer" onClick={handleClick}>
      <div className="relative">
        <AspectRatio ratio={4/3}>
          <img 
            src={imageUrl} 
            alt={title} 
            className="object-cover w-full h-full"
          />
        </AspectRatio>
        <div className="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent"></div>
        <div className="absolute bottom-0 left-0 right-0 p-4 text-white">
          <h3 className="font-semibold text-lg">{title}</h3>
          <div className="flex items-center text-sm">
            <Map size={14} className="mr-1" />
            <span>{location}</span>
          </div>
        </div>
      </div>
      <div className="p-3 flex justify-between items-center">
        <div className="text-sm text-gray-600">{days} วัน</div>
        <button className="text-xs apple-button px-3 py-1 rounded-full text-white">
          เลือก
        </button>
      </div>
    </div>
  );
};

const RecommendedTrips = () => {
  const trips = [
    {
      id: 1,
      title: "เที่ยวกรุงเทพฯ 3 วัน",
      location: "กรุงเทพมหานคร",
      imageUrl: "https://images.unsplash.com/photo-1601192115542-5feb5efb9b5c",
      days: 3,
    },
    {
      id: 2,
      title: "เชียงใหม่สุดชิล",
      location: "เชียงใหม่",
      imageUrl: "https://images.unsplash.com/photo-1512553353614-82a7370096dc",
      days: 4,
    },
    {
      id: 3,
      title: "ทะเลภูเก็ต",
      location: "ภูเก็ต",
      imageUrl: "https://images.unsplash.com/photo-1589394815804-964ed0be2eb5",
      days: 5,
    },
    {
      id: 4,
      title: "เกาะช้างสุดฮิต",
      location: "ตราด",
      imageUrl: "https://images.unsplash.com/photo-1552465011-b4e21bf6e79a",
      days: 3,
    },
  ];

  return (
    <div className="p-4">
      <h3 className="text-lg font-medium mb-3">ทริปแนะนำ</h3>
      <Carousel>
        <CarouselContent>
          {trips.map((trip) => (
            <CarouselItem key={trip.id} className="md:basis-1/2 lg:basis-1/2">
              <TripCard
                title={trip.title}
                location={trip.location}
                imageUrl={trip.imageUrl}
                days={trip.days}
              />
            </CarouselItem>
          ))}
        </CarouselContent>
      </Carousel>
    </div>
  );
};

export default RecommendedTrips;
