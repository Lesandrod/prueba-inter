function isRectangular(matrix) {
  if (!Array.isArray(matrix) || matrix.length === 0) return false;

  const cols = matrix[0].length;

  return matrix.every((row) => row.length === cols);
}

function isDiagonal(matrix) {
  if (matrix.length !== matrix[0].length) return false;

  for (let i = 0; i < matrix.length; i++) {
    for (let j = 0; j < matrix[i].length; j++) {
      if (i !== j && matrix[i][j] !== 0) {
        return false;
      }
    }
  }

  return true;
}

function computeStats({ q, r }) {
  const values = [...q.flat(), ...r.flat()];

  const total = values.reduce((sum, value) => sum + value, 0);

  return {
    max: Math.max(...values),
    min: Math.min(...values),
    average: total / values.length,
    total,
    hasDiagonalMatrix: isDiagonal(q) || isDiagonal(r),
  };
}

module.exports = {
  computeStats,
};
