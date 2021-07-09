import { useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import classNames from 'classnames';
import { useAtom } from 'jotai';
import { searchFocusAtom } from '../store/store';

export const Layout = ({ children }) => {
  let location = useLocation();
  const [isSearchFocus] = useAtom(searchFocusAtom);

  useEffect(() => {
    window.scrollTo({
      top: 0,
      behavior: 'smooth',
    });
  }, [location]);

  return (
    <>
      <div className="w-full h-full pt-28">
        {/*Overlay on search focus*/}
        <div
          className={classNames(
            {
              'bg-opacity-40': isSearchFocus,
              'bg-opacity-0': !isSearchFocus,
            },
            'fixed w-full h-full bg-black z-20 pointer-events-none transition-all'
          )}
        />

        {/*Content*/}
        {children}
      </div>
    </>
  );
};
