/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./content/**/*.{html,js}", "./layouts/**/*.{html,js}"],
  theme: {
    colors: {
      transparent: "transparent",
      current: "currentColor",
      white: "#f8f8f2",
      blue: "#7690e3",
      "light-black": "#21242e",
      black: "#14161c",
      blacker: "#000000",
      "light-gray": "#b3b3b5",
      gray: "#282a36",
      "dark-gray": "#1a1c24",
      "accent-color": "#30324d",
    },
  },
  plugins: [],
};
