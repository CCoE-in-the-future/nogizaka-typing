import React from 'react';
import Header from '../presentation/components/Header'; // Header コンポーネントをインポート
import HomeView from '../presentation/components/HomeView';

const Home: React.FC = () => {
  return (
    <div>
      <Header /> 
      <main className="px-4 py-8">
        <HomeView />
      </main>
    </div>
  );
};

export default Home;
