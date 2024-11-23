// src/app/components/Header.tsx

import React from 'react';
import Image from 'next/image'

const HomeView: React.FC = () => {
  return (
    <div id="main" className="flex flex-col items-center">
        <Image
          className="dark:invert"
          src="/aruno.jpeg"
          alt="Next.js logo"
          width={180}
          height={38}
          priority
        />
    </div>
  );
};

export default HomeView;
