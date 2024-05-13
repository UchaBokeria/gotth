/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    'server/view/**/*.templ',
  ],
  theme: {
    fontFamily: {
      nino: ["Nino", "sans-serif"],
      deja: ["Deja", "sans-serif"],
      arial: ["Arial", "sans-serif"],
      poppins: ["Poppins", "sans-serif"],
    },
    container: {
      center: true,
      padding: {
        DEFAULT: "1rem",
        mobile: "2rem",
        tablet: "4rem",
        desktop: "5rem",
      },
    },
    extend: {
      screens: {
        'mobile': '1319px',
      },
      colors: {
        primary: "#008772",
        secondary: "#FFF200",
        white: "#FFFFFF",
        black: "#333333",
      }
    },
  },
  plugins: []
}