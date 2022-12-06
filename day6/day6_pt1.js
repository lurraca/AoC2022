const fs = require('fs');

const markerSize = 4;

fs.readFile('input.txt', 'utf8', (err, data) => {
  if (err) {
    console.error(err);
    return;
  }
  solution = findSOPMarker(data);
  console.log(solution);
});


function findSOPMarker(data) {
  let marker = 0;
  let map1 = new Map();
  let inputAr = data.split("");

  while(map1.size < markerSize) {
    for (const [index, char] of inputAr.entries()) {

      if (map1.size == markerSize)
        break;

      map1.set(char, true);

      for (const el of inputAr.slice(index+1)) {

        if (map1.size == markerSize) {
          marker = index + markerSize;
          break;
        }

        if (map1.get(el)) {
          map1.clear();
          break;
        } else {
          map1.set(el, true);
        }
      }
    }
  }

  return marker;
}

