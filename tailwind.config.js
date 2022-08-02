/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./layouts/**/*",
    "./docs/**/*.html",
    "./content/**/*.md",
  ],
  theme: {
    extend: {},
    colors: {
      transparent: 'transparent',
      current: 'currentColor',
      'white': '#f8f8f2',
      'blue': '#6483e6',
      'light-black': '#23242b',
      'black': '#18191f',
      'blacker': '#000000',
      'light-gray': '#b3b3b5',
      'gray': '#282a36',
      'dark-gray': '#14151b',
      'accent-color': '#383a59'
    },
  },
  plugins: [],
}
