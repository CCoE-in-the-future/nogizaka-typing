// src/app/components/Header.tsx

import React from 'react';

const Header: React.FC = () => {
  return (
    <header className="bg-black shadow-md">
      <nav className="max-w-screen-xl mx-auto px-4 py-3 flex items-center justify-between">
        <a href="/pages" className="text-xl font-semibold text-white">Nogizaka Typing</a>
        <div className="flex space-x-4">
          {/* <a href="/pages" className="text-white hover:text-blue-600">Home</a> */}
          <a href="/signin" className="text-white hover:text-blue-600">Sign in</a>
        </div>
      </nav>
    </header>
  );
};

export default Header;
