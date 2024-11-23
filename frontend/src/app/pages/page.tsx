// src/pages/index.tsx

import React from 'react';
import Header from '../components/Header'; // Header コンポーネントをインポート
import HomeView from '../components/HomeView';

const Home: React.FC = () => {
  return (
    <div>
      <Header /> {/* Header コンポーネントを表示 */}
      <main className="px-4 py-8">
        <HomeView />
      </main>
    </div>
  );
};

export default Home;
