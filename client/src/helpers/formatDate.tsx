import React from "react";

function formatDate(date: Date) {
  return date.toString().slice(0, 15);
}

export default formatDate;
