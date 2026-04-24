/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  corePlugins: {
    preflight: false, // THIS IS THE KEY
  },
  theme: {
    extend: {},
  },
  plugins: [],
}