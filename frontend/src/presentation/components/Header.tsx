import Link from 'next/link';
import React from 'react';

const Header: React.FC = () => {
  return (
    <header className="bg-black shadow-md">
      <nav className="max-w-screen-xl mx-auto px-4 py-3 flex items-center justify-between">
        <Link href="/" className="text-xl font-semibold text-white">Nogizaka Typing</Link>
        <div className="flex space-x-4">
          <Link href="/signin" className="text-white hover:text-blue-600">Sign in</Link>
        </div>
      </nav>
    </header>
  );
};

export default Header;
