import React from 'react';
import Header from '../../presentation/components/Header'; // Header コンポーネントをインポート
import SelectTarget from './components/SelectTarget';

const Home: React.FC = () => {
  return (
    <div>
      <Header /> {/* Header コンポーネントを表示 */}
      <main className="px-4 py-8">
        <SelectTarget />
      </main>
    </div>
  );
};

export default Home;
