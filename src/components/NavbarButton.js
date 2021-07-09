export const NavbarButton = ({ children }) => {
  return (
    <button
      type="button"
      className="py-1 px-3 rounded-md text-xs hover:bg-gray-100 hover:text-xgreen transition-all"
    >
      {children}
    </button>
  );
};
