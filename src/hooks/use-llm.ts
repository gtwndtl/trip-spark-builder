
import { useState } from 'react';
import { useToast } from '@/hooks/use-toast';

export interface LLMResponse {
  parsed?: {
    destination?: string;
    days?: number;
    style?: string;
    budget?: number;
  };
  error?: string;
}

interface UseLLMOptions {
  endpoint?: string;
  timeout?: number;
}

export function useLLM(options: UseLLMOptions = {}) {
  const { endpoint = 'http://127.0.0.1:8000', timeout = 30000 } = options;
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const { toast } = useToast();

  const sendMessage = async (message: string): Promise<LLMResponse> => {
    setLoading(true);
    setError(null);
    
    // Create a timeout promise
    const timeoutPromise = new Promise<never>((_, reject) => {
      setTimeout(() => reject(new Error('Request timed out')), timeout);
    });
    
    try {
      // Race the fetch against the timeout
      const response = await Promise.race([
        fetch(`${endpoint}/plan-trip`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ message }),
        }),
        timeoutPromise,
      ]) as Response;
      
      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(`API Error: ${response.status} - ${errorText || response.statusText}`);
      }
      
      const data = await response.json();
      return data;
    } catch (error) {
      const errorMessage = error instanceof Error ? error.message : 'Unknown error occurred';
      setError(errorMessage);
      
      toast({
        title: "เกิดข้อผิดพลาด",
        description: errorMessage,
        variant: "destructive",
      });
      
      return { error: errorMessage };
    } finally {
      setLoading(false);
    }
  };
  
  const planTrip = async (preferences: any): Promise<any> => {
    setLoading(true);
    setError(null);
    
    try {
      const response = await fetch(`${endpoint}/api/trip_planner`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ preferences }),
      });
      
      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(`API Error: ${response.status} - ${errorText || response.statusText}`);
      }
      
      return await response.json();
    } catch (error) {
      const errorMessage = error instanceof Error ? error.message : 'Unknown error occurred';
      setError(errorMessage);
      
      toast({
        title: "เกิดข้อผิดพลาด",
        description: errorMessage,
        variant: "destructive",
      });
      
      return { error: errorMessage };
    } finally {
      setLoading(false);
    }
  };

  return {
    sendMessage,
    planTrip,
    loading,
    error,
  };
}
