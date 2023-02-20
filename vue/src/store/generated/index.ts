// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import MoonIbank from './moon.ibank'
import MoonMoon from './moon.moon'


export default { 
  MoonIbank: load(MoonIbank, 'moon.ibank'),
  MoonMoon: load(MoonMoon, 'moon.moon'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}