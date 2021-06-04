var singletonMarket = (function () {
 
    // armazena uma referência ao singleton
    var instance;
 
    function init(){
    
        // métodos e propriedades do singleton 
 
        // objetos privados
        var cart = [];
        var deleteItem = function (item) {
            index = cart.indexOf(item);
            if (index > -1) {
                cart.splice(index, 1);
            }
        };

        // objetos públicos
        return {
            list: cart,
            add: function (item) {
                cart.push(item);
            },
            delete: function (item) {
                deleteItem(item);
            }
        };
    };

    // retorna sempre a mesma instância dessa classe
    return {
 
        // se a instância existir, retorna ela, senão cria uma nova
        getInstance: function () {
 
            if ( !instance ) {
                instance = init();
            }
 
            return instance;
        }
    };
})();

function run(){
    est1 = singletonMarket.getInstance();
    est2 = singletonMarket.getInstance();
    alert(`Única instância = ${est1===est2}`);
}