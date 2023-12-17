/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/assets/templates/**/*.{html,js}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["light", "dark", "dim", "nord", "winter", "corporate"],
  },
}
