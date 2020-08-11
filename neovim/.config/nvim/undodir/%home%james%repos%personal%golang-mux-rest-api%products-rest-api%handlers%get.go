Vim�UnDo� գ���[�C�J�Y�1�F��4���2�J�;47:   C                                   _2c}    _�                             ����                                                                                                                                                                                                                                                                                                                                                             _2co     �         C    5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             _2cp     �         D    �         D    5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             _2cq     �                -	"github/JamieBShaw/golang-mux-rest-api/data"5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             _2cr     �                -	"github/JamieBShaw/golang-mux-rest-api/data"5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             _2c{     �         C    �         C    5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             _2c|     �         D       5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             _2c|    �               C   package handlers       import (   ?	"github/JamieBShaw/golang-mux-rest-api/products-rest-api/data"   	"net/http"   )       4// swagger:route GET /products products listProducts   /// Returns a list of products from the database   // responses:   //  200: productsResponse       @// ListAll handles GET requests and returns all current products   Efunc (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {   '	p.l.Println("[DEBUG] get all records")       4	rw.Header().Add("Content-Type", "application/json")       )	// fetch the products from the datastore   	prods := data.GetProducts()       	// serialize the list to JSON   	err := data.ToJSON(prods, rw)   	if err != nil {   J		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)   	}   }       >// swagger:route GET /products/{id} products listSingleProduct   -// Returns a single product from the database   // responses:   //  200: productsResponse   //  404: errorResponse       "// ListSingle handles GET requests   Hfunc (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {       4	rw.Header().Add("Content-Type", "application/json")   	id := getProductID(r)       )	p.l.Println("[DEBUG] get record id", id)       %	prod, err := data.GetProductByID(id)       	switch err {   
	case nil:       	case data.ErrProductNotFound:   .		p.l.Println("[ERROR] fetching product", err)       %		rw.WriteHeader(http.StatusNotFound)   6		data.ToJSON(&GenericError{Message: err.Error()}, rw)   		return   		default:   .		p.l.Println("[ERROR] fetching product", err)       0		rw.WriteHeader(http.StatusInternalServerError)   6		data.ToJSON(&GenericError{Message: err.Error()}, rw)   		return   	}       	err = data.ToJSON(prod, rw)   	if err != nil {   :		// we should never be here but log the error just incase   1		p.l.Println("[ERROR] serializing product", err)   	}   }5��