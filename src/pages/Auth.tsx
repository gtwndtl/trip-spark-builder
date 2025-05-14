
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Header from '@/components/Header';
import Footer from '@/components/Footer';
import LoginForm from '@/components/auth/LoginForm';
import RegisterForm from '@/components/auth/RegisterForm';
import { useToast } from '@/hooks/use-toast';

const Auth = () => {
  const [activeTab, setActiveTab] = useState<'login' | 'register'>('login');
  const navigate = useNavigate();
  const { toast } = useToast();

  const handleSuccessfulAuth = () => {
    toast({
      title: activeTab === 'login' ? "เข้าสู่ระบบสำเร็จ" : "ลงทะเบียนสำเร็จ",
      description: "ยินดีต้อนรับสู่ TRIP PLANNER",
    });
    navigate('/');
  };

  return (
    <div className="min-h-screen flex flex-col bg-gray-50">
      <Header />
      
      <main className="flex-1 flex items-center justify-center px-4 py-12">
        <div className="glass-card w-full max-w-md p-6 rounded-xl">
          <div className="mb-8">
            <h2 className="text-3xl text-center font-light">
              <span className="text-tripPurple font-normal">TRIP</span> PLANNER
            </h2>
            <p className="text-center text-gray-600 mt-2">
              {activeTab === 'login' ? 'เข้าสู่ระบบเพื่อจัดการทริปของคุณ' : 'สร้างบัญชีเพื่อเริ่มต้นการวางแผนทริป'}
            </p>
          </div>

          <div className="flex mb-6 border-b">
            <button 
              className={`flex-1 py-3 text-center transition-colors ${activeTab === 'login' ? 'text-tripPurple border-b-2 border-tripPurple font-medium' : 'text-gray-500 hover:text-tripPurple'}`}
              onClick={() => setActiveTab('login')}
            >
              เข้าสู่ระบบ
            </button>
            <button 
              className={`flex-1 py-3 text-center transition-colors ${activeTab === 'register' ? 'text-tripPurple border-b-2 border-tripPurple font-medium' : 'text-gray-500 hover:text-tripPurple'}`}
              onClick={() => setActiveTab('register')}
            >
              ลงทะเบียน
            </button>
          </div>

          {activeTab === 'login' ? (
            <LoginForm onSuccess={handleSuccessfulAuth} />
          ) : (
            <RegisterForm onSuccess={handleSuccessfulAuth} />
          )}

          <div className="mt-8 text-center text-sm text-gray-500">
            {activeTab === 'login' ? (
              <p>
                ยังไม่มีบัญชี?{' '}
                <button 
                  onClick={() => setActiveTab('register')} 
                  className="text-tripPurple hover:underline"
                >
                  ลงทะเบียนที่นี่
                </button>
              </p>
            ) : (
              <p>
                มีบัญชีอยู่แล้ว?{' '}
                <button 
                  onClick={() => setActiveTab('login')} 
                  className="text-tripPurple hover:underline"
                >
                  เข้าสู่ระบบที่นี่
                </button>
              </p>
            )}
          </div>
        </div>
      </main>
      
      <Footer />
    </div>
  );
};

export default Auth;
