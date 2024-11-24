// src/app/components/Header.tsx

import React from 'react';
import Image from 'next/image'
import Link from 'next/link';

const HomeView: React.FC = () => {
  return (
    <div id="main" className="flex flex-col items-center">
      <Image
        className="dark:invert"
        src="/aruno.jpeg"
        alt="Next.js logo"
        width={300}
        height={200}
        priority
      />
      <Link href="/select-target" className="bg-black text-white px-4 py-2 rounded border border-white hover:bg-gray-800 active:bg-gray-700 transition-colors duration-200">
        Select Target
      </Link>
    </div>
  );
};

export default HomeView;
