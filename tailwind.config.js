const colors = require("tailwindcss/colors");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "internal/templates/*.templ",
    "internal/templates/components/*.templ",
  ],
  theme: {
    extend: {
      colors: {
        "floral-white": "#FFFAF0",
        "soft-blue": "#A0C4FF",
        coral: "#FF7F50",
        "dark-navy": "#1A2A40",
        "muted-blue": "#4A6FA5",
        "muted-coral": "#E76F51",
        "vd-blue": "#0F172A",
      },
      backgroundColor: {
        light: "#FFFAF0", // floral-white
        dark: "#0F172A", // vd-blue
      },
      textColor: {
        dim: "#1A2A40", // dark-navy
        lite: "#FFFAF0", // floral-white
      },
    },
  },
  darkMode: "class",
  plugins: [
    require("@tailwindcss/forms"),
    require("@tailwindcss/typography"),
    function ({ addUtilities }) {
      const newUtilities = {
        ".details-reset": {
          "& summary": {
            "list-style": "none",
            "&::-webkit-details-marker": {
              display: "none",
            },
          },
        },
      };
      addUtilities(newUtilities, ["responsive", "hover"]);
    },
  ],
};
