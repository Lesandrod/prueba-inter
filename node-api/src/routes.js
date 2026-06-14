const jwt = require("jsonwebtoken");
const { computeStats } = require("../funtions/stats");

const SECRET = "secret";

function auth(req, res, next) {
  const token = req.headers.authorization?.split(" ")[1];

  if (!token) {
    return res.status(401).json({
      error: "token requerido",
    });
  }

  try {
    jwt.verify(token, SECRET);
    next();
  } catch {
    res.status(401).json({
      error: "token invalido",
    });
  }
}

module.exports = function (app) {
  app.post("/login", (req, res) => {
    const { username, password } = req.body;

    if (username !== "admin123" || password !== "prueba123") {
      return res.status(401).json({
        error: "credenciales invalidas",
      });
    }

    const token = jwt.sign({ username }, SECRET, { expiresIn: "1h" });

    res.json({ token });
  });

  app.post("/stats", auth, (req, res) => {
    try {
      const stats = computeStats(req.body);

      res.json(stats);
    } catch (error) {
      res.status(400).json({
        error: error.message,
      });
    }
  });
};
