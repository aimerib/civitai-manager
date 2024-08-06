module.exports = {
  content: [
    "./templates/**/*.html",
    "./templates/**/*.plush.html",
    "./assets/js/**/*.js",
  ],
  darkMode: "class",
  theme: {
    extend: {
      colors: {
        primary: {
          100: "#e6f0ff",
          200: "#cce0ff",
          300: "#63b3ed",
          400: "#4299e1",
          500: "#0066ff",
          600: "#0052cc",
          700: "#003d99",
          800: "#002966",
          900: "#001433",
        },
        secondary: {
          100: "#fff0e6",
          200: "#ffe0cc",
          300: "#ffc299",
          400: "#ed8936",
          500: "#ff6600",
          600: "#cc5200",
          700: "#993d00",
          800: "#662900",
          900: "#331400",
        },
        background: {
          dark: "#1a202c",
          card: "#2d3748",
        },
        text: {
          light: "#ffffff",
          muted: "#909296",
        },
        border: "#373a40",
      },
      fontFamily: {
        sans: [
          "Inter",
          "-apple-system",
          "BlinkMacSystemFont",
          "Segoe UI",
          "Roboto",
          "Oxygen",
          "Ubuntu",
          "Cantarell",
          "Fira Sans",
          "Droid Sans",
          "Helvetica Neue",
          "sans-serif",
        ],
      },
      fontSize: {
        sm: "12px",
        base: "14px",
        lg: "16px",
      },
      fontWeight: {
        normal: 400,
        bold: 700,
      },
      lineHeight: {
        base: 1.5,
      },
      spacing: {
        xs: "4px",
        sm: "8px",
        md: "16px",
        lg: "24px",
        xl: "32px",
      },
      borderRadius: {
        sm: "4px",
        md: "8px",
        lg: "12px",
      },
      boxShadow: {
        card: "0 1px 3px rgba(0, 0, 0, 0.1)",
        dropdown: "0 4px 6px rgba(0, 0, 0, 0.1)",
      },
      height: {
        header: "60px",
      },
      width: {
        sidebar: "240px",
      },
      maxWidth: {
        content: "1200px",
      },
      zIndex: {
        header: 1000,
        dropdown: 1050,
        modal: 1100,
      },
      transitionDuration: {
        fast: "200ms",
        medium: "300ms",
      },
      transitionTimingFunction: {
        ease: "ease",
      },
    },
  },
  plugins: [],
}
