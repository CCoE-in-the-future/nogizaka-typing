import React from 'react';
import Image from 'next/image';
import Link from 'next/link';

const SelectTarget: React.FC = () => {
  return (
    <div className="flex flex-col items-center justify-center">
      <h1 className="text-4xl font-bold mb-6 text-gray-800 dark:text-white">
        Select Generations
      </h1>

      <div id="main" className="flex flex-row items-center space-x-4">
        {[['1st', '/nanase.jpeg'], ['2nd', '/miona.jpeg'], ['3rd', '/yuki.jpeg'], ['4th', '/haruka.jpeg'], ['5th', '/aruno.jpeg']].map(([generation, img]) => (
          <div key={generation} className="flex flex-col items-center">
            <Image
              className="dark:invert"
              src={img}
              alt={`${generation} image`}
              width={200}
              height={100}
              priority
            />
            <div className="flex items-center space-x-2 mt-2">
              <input type="checkbox" className="w-5 h-5 accent-green-500" />
              <span>{generation}</span>
            </div>
          </div>
        ))}
      </div>

      <div className="mt-8">
        <Link
          href="/"
          className="bg-gray-600 text-white px-6 py-3 rounded border border-white hover:bg-gray-700 active:bg-gray-800 transition-colors duration-200">
          Return home
        </Link>
      </div>
    </div>
  );
};

export default SelectTarget;
