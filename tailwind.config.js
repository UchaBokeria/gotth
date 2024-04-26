/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    'server/view/**/*.templ',
  ],
  theme: {
    fontFamily: {
      poppins: ["Nino", "sans-serif"],
      poppins: ["Deja", "sans-serif"],
      poppins: ["Arial", "sans-serif"],
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