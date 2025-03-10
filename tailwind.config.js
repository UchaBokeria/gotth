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
        // 'md': { 'raw': '(min-height: 800px)' },
        'mob':  {'min': '1px', 'max': '1280px'},
        
        // 'md': '1280px', // Adjust the value as needed
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