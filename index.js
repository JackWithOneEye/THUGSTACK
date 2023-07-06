import htmx from 'htmx.org';

window.htmx = htmx;

htmx.defineExtension('checkbox-enc', {
  encodeParameters: (xhr, parameters, elt) => {
    const cbs = elt.querySelectorAll('input[type=checkbox]');
    cbs.forEach((cb) => {
      if (cb.name) {
        parameters[cb.name] = cb.checked === true;
      }
    });
  }
});