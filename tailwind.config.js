module.exports = {
  purge: ['./src/**/*.{js,jsx,ts,tsx}', './public/index.html'],
  theme: {
    fontFamily: {
      sans: [
        'Open Sans',
        '-apple-system',
        'sans-serif',
        'Apple Color Emoji',
        'Segoe UI Emoji',
        'Segoe UI Symbol',
        'Noto Color Emoji',
      ],
    },
    extend: {
      colors: {
        xgreen: '#03AC0E',
        xgreendark: '#12883D',
      },
    },
  },
  variants: {
    extend: {
      textColor: ['disabled'],
      opacity: ['disabled'],
      pointerEvents: ['disabled'],
    },
  },
  plugins: [],
};
