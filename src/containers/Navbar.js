import { Link } from 'react-router-dom';
import { useAtom } from 'jotai';

import {
  searchFocusAtom,
  postCodeAtom,
  sortAtom,
  queryAtom,
} from '../store/store';

import { NavbarButton } from '../components';

export const Navbar = () => {
  const [, setSearchFocus] = useAtom(searchFocusAtom);
  const [postCode, setPostCode] = useAtom(postCodeAtom);
  const [query, setQuery] = useAtom(queryAtom);
  const [sort] = useAtom(sortAtom);

  return (
    <nav className="fixed top-0 w-full md:h-50 z-50 flex flex-col content-center justify-center py-2 bg-white shadow-lg">
      {/*Top thin bar*/}
      <div className="self-center flex flex-row flex-grow w-full px-8 pb-1 content-center justify-between bg-gray-100 text-gray-400">
        <div>
          <button className="flex flex-row items-center space-x-1">
            <i className="icon-navbar phone" />
            <span className="text-xs hover:text-xgreen">
              Download Tokopedia App
            </span>
          </button>
        </div>
        <div className="flex flex-row space-x-4">
          <button className="flex flex-row items-center space-x-1">
            <span className="text-xs hover:text-xgreen">Tentang Tokopedia</span>
          </button>
          <button className="flex flex-row items-center space-x-1">
            <span className="text-xs hover:text-xgreen">Mitra Tokopedia</span>
          </button>
          <button className="flex flex-row items-center space-x-1">
            <span className="text-xs hover:text-xgreen">
              Pusat Edukasi Seller
            </span>
          </button>
          <button className="flex flex-row items-center space-x-1">
            <span className="text-xs hover:text-xgreen">Promo</span>
          </button>
          <button className="flex flex-row items-center space-x-1">
            <span className="text-xs hover:text-xgreen">Tokopedia Care</span>
          </button>
        </div>
      </div>

      {/*Main bar*/}
      <div className="self-center flex flex-row flex-grow px-10 py-3 content-center justify-between w-full bg-white text-gray-400">
        <div className="flex flex-row space-x-4">
          {/*Logo Tokopedia*/}
          <Link to="/" className="flex self-center text-md w-36">
            <img className="" src="/assets/logo-tokopedia.svg" alt="logo" />
          </Link>

          {/*Kategori button*/}
          <NavbarButton>Kategori</NavbarButton>
        </div>

        {/*Search bar*/}
        <form
          action="/search"
          method="get"
          className="flex flex-row w-full h-8 mx-4 border rounded-lg border-gray-300 focus-within:border-xgreen focus-within:outline-none focus-within:ring-0 focus-within:shadow-outline transition-all"
        >
          <input
            className="w-full rounded-l-lg px-3 py-1.5 text-xs text-gray-800 leading-5 focus:outline-none"
            placeholder="Cari di Tokopedia"
            name="q"
            onFocus={() => setSearchFocus(true)}
            onBlur={() => setSearchFocus(false)}
            onChange={(e) => setQuery(e.target.value)}
            value={query}
            type="text"
          />
          <input name="sort" value={sort} hidden />
          <input name="from" value={postCode} hidden />
          <button
            aria-label="Tombol pencarian"
            type="submit"
            className="flex rounded-r-lg h-full w-10 bg-gray-100 items-center justify-center"
          >
            <i className="icon-search" />
          </button>
        </form>

        {/*Misc menus*/}
        <div className="flex flex-row divide-x divide-gray-300">
          <div className="flex flex-row px-4 space-x-4">
            <NavbarButton>
              <i className="icon-navbar cart" />
            </NavbarButton>
            <NavbarButton>
              <i className="icon-navbar bell" />
            </NavbarButton>
            <NavbarButton>
              <i className="icon-navbar mail" />
            </NavbarButton>
          </div>
          <div className="flex flex-row px-4 space-x-16">
            <button
              type="button"
              className="flex flex-row py-1 px-3 space-x-2 content-center justify-center rounded-md text-xs hover:bg-gray-100 hover:text-xgreen"
            >
              <img
                className="self-center w-6 h-6 rounded-full"
                src="/assets/img-toko.png"
              />
              <div className="self-center">Toko</div>
            </button>
            <button
              type="button"
              className="flex flex-row py-1 px-3 space-x-2 content-center justify-center rounded-md text-xs hover:bg-gray-100 hover:text-xgreen"
            >
              <img
                className="self-center w-6 h-6 rounded-full"
                src="/assets/img-user.png"
              />
              <div className="self-center">User</div>
            </button>
          </div>
        </div>
      </div>

      <div className="self-center flex flex-row flex-grow w-full px-12 content-center justify-end text-gray-400">
        <button className="flex flex-row items-center space-x-1">
          <svg
            className="unf-icon"
            viewBox="0 0 24 24"
            width="14"
            height="14"
            fill="var(--color-icon-enabled, #525867)"
            style={{
              display: 'inline-block',
              verticalAlign: 'middle',
            }}
          >
            <path
              fillRule="evenodd"
              clipRule="evenodd"
              d="M11.5 21.87a1 1 0 00.5.13 1 1 0 00.514-.138C12.974 21.587 20 17.399 20 10a7.909 7.909 0 00-8-8 7.91 7.91 0 00-8 8c0 7.399 7.025 11.587 7.486 11.862l.014.008zM9.694 4.44A5.94 5.94 0 0112 4a5.94 5.94 0 016 6c0 5.28-4.48 8.81-6 9.81-1.52-1.03-6-4.51-6-9.81a5.94 5.94 0 013.694-5.56zm.084 8.886a4 4 0 104.444-6.652 4 4 0 00-4.444 6.652zm1.11-4.989a2 2 0 112.223 3.326 2 2 0 01-2.222-3.326z"
            />
          </svg>
          <button
            onClick={() => {
              postCode === '17612'
                ? setPostCode('36138')
                : setPostCode('17612');
            }}
            className="flex flex-row space-x-1 items-center text-xs text-gray-700 hover:text-xgreen"
          >
            <span>Dikirim ke</span>
            <span className="font-bold">
              {postCode === '17612'
                ? 'Bekasi, Jawa Barat'
                : 'Kota Jambi, Jambi'}
            </span>
            <i className="icon-navbar arrowdown" />
          </button>
        </button>
      </div>
    </nav>
  );
};

export default Navbar;
