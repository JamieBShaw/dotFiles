Vim�UnDo� ����y���-AuPS!��� ��Y��6%�K   U                                   _5L    _�                             ����                                                                                                                                                                                                                                                                                                                                                             _2c�     �         L    �         L    5�_�                           ����                                                                                                                                                                                                                                                                                                                                         3       v���    _2c�     �         M      3import "github/JamieBShaw/golang-mux-rest-api/data"   ?	"github/JamieBShaw/golang-mux-rest-api/products-rest-api/data"5�_�                           ����                                                                                                                                                                                                                                                                                                                                         3       v���    _2c�     �         L      Fimport 	"github/JamieBShaw/golang-mux-rest-api/products-rest-api/data"5�_�                           ����                                                                                                                                                                                                                                                                                                                                         3       v���    _2c�    �               L   (// Package classification of Product API   //    // Documentation for Product API   //   //	Schemes: http   //	BasePath: /   //	Version: 1.0.0   //   //	Consumes:   //	- application/json   //   //	Produces:   //	- application/json   //   // swagger:meta   package handlers       Eimport "github/JamieBShaw/golang-mux-rest-api/products-rest-api/data"       //   A// NOTE: Types defined here are purely for documentation purposes   2// these types are not used by any of the handlers       -// Generic error message returned as a string   !// swagger:response errorResponse   "type errorResponseWrapper struct {   	// Description of the error   	// in: body   	Body GenericError   }       3// Validation errors defined as an array of strings   #// swagger:response errorValidation   $type errorValidationWrapper struct {   	// Collection of the errors   	// in: body   	Body ValidationError   }       // A list of products   $// swagger:response productsResponse   %type productsResponseWrapper struct {   	// All current products   	// in: body   	Body []data.Product   }       /// Data structure representing a single product   #// swagger:response productResponse   $type productResponseWrapper struct {   	// Newly created product   	// in: body   	Body data.Product   }       .// No content is returned by this API endpoint   %// swagger:response noContentResponse   &type noContentResponseWrapper struct {   }       1// swagger:parameters updateProduct createProduct   "type productParamsWrapper struct {   /	// Product data structure to Update or Create.   A	// Note: the id field is ignored by update and create operations   	// in: body   	// required: true   	Body data.Product   }       5// swagger:parameters listSingleProduct deleteProduct   $type productIDParamsWrapper struct {   9	// The id of the product for which the operation relates   	// in: path   	// required: true   	ID int `json:"id"`   }5�_�                           ����                                                                                                                                                                                                                                                                                                                                         3       v���    _2c�    �               L   (// Package classification of Product API   //    // Documentation for Product API   //   //	Schemes: http   //	BasePath: /   //	Version: 1.0.0   //   //	Consumes:   //	- application/json   //   //	Produces:   //	- application/json   //   // swagger:meta   package handlers       Eimport "github/JamieBShaw/golang-mux-rest-api/products-rest-api/data"       //   A// NOTE: Types defined here are purely for documentation purposes   2// these types are not used by any of the handlers       -// Generic error message returned as a string   !// swagger:response errorResponse   "type errorResponseWrapper struct {   	// Description of the error   	// in: body   	Body GenericError   }       3// Validation errors defined as an array of strings   #// swagger:response errorValidation   $type errorValidationWrapper struct {   	// Collection of the errors   	// in: body   	Body ValidationError   }       // A list of products   $// swagger:response productsResponse   %type productsResponseWrapper struct {   	// All current products   	// in: body   	Body []data.Product   }       /// Data structure representing a single product   #// swagger:response productResponse   $type productResponseWrapper struct {   	// Newly created product   	// in: body   	Body data.Product   }       .// No content is returned by this API endpoint   %// swagger:response noContentResponse   &type noContentResponseWrapper struct {   }       1// swagger:parameters updateProduct createProduct   "type productParamsWrapper struct {   /	// Product data structure to Update or Create.   A	// Note: the id field is ignored by update and create operations   	// in: body   	// required: true   	Body data.Product   }       5// swagger:parameters listSingleProduct deleteProduct   $type productIDParamsWrapper struct {   9	// The id of the product for which the operation relates   	// in: path   	// required: true   	ID int `json:"id"`   }5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             _4_�     �                Eimport "github/JamieBShaw/golang-mux-rest-api/products-rest-api/data"5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             _4_�    �               L   (// Package classification of Product API   //    // Documentation for Product API   //   //	Schemes: http   //	BasePath: /   //	Version: 1.0.0   //   //	Consumes:   //	- application/json   //   //	Produces:   //	- application/json   //   // swagger:meta   package handlers       Iimport "github.com/JamieBShaw/golang-mux-rest-api/products-rest-api/data"       //   A// NOTE: Types defined here are purely for documentation purposes   2// these types are not used by any of the handlers       -// Generic error message returned as a string   !// swagger:response errorResponse   "type errorResponseWrapper struct {   	// Description of the error   	// in: body   	Body GenericError   }       3// Validation errors defined as an array of strings   #// swagger:response errorValidation   $type errorValidationWrapper struct {   	// Collection of the errors   	// in: body   	Body ValidationError   }       // A list of products   $// swagger:response productsResponse   %type productsResponseWrapper struct {   	// All current products   	// in: body   	Body []data.Product   }       /// Data structure representing a single product   #// swagger:response productResponse   $type productResponseWrapper struct {   	// Newly created product   	// in: body   	Body data.Product   }       .// No content is returned by this API endpoint   %// swagger:response noContentResponse   &type noContentResponseWrapper struct {   }       1// swagger:parameters updateProduct createProduct   "type productParamsWrapper struct {   /	// Product data structure to Update or Create.   A	// Note: the id field is ignored by update and create operations   	// in: body   	// required: true   	Body data.Product   }       5// swagger:parameters listSingleProduct deleteProduct   $type productIDParamsWrapper struct {   9	// The id of the product for which the operation relates   	// in: path   	// required: true   	ID int `json:"id"`   }5�_�      	              D        ����                                                                                                                                                                                                                                                                                                                                                             _5~�     �   D   K   M       �   D   F   L    5�_�      
           	   I        ����                                                                                                                                                                                                                                                                                                                                                             _5~�     �   H   K   R       5�_�   	              
   F        ����                                                                                                                                                                                                                                                                                                                                                             _5~�     �   F   I   T      	�   F   H   S    5�_�   
                 I       ����                                                                                                                                                                                                                                                                                                                                                             _5     �   H   J   W      	//in: query5�_�                    J       ����                                                                                                                                                                                                                                                                                                                                                             _5     �   I   K   W      	//5�_�                    M        ����                                                                                                                                                                                                                                                                                                                                                             _5	     �   L   M           5�_�                    L        ����                                                                                                                                                                                                                                                                                                                                                             _5	     �   K   L           5�_�                    K        ����                                                                                                                                                                                                                                                                                                                                                             _5
     �   J   K           5�_�                    K       ����                                                                                                                                                                                                                                                                                                                                                             _5    �               T   (// Package classification of Product API   //    // Documentation for Product API   //   //	Schemes: http   //	BasePath: /   //	Version: 1.0.0   //   //	Consumes:   //	- application/json   //   //	Produces:   //	- application/json   //   // swagger:meta   package handlers       Iimport "github.com/JamieBShaw/golang-mux-rest-api/products-rest-api/data"       //   A// NOTE: Types defined here are purely for documentation purposes   2// these types are not used by any of the handlers       -// Generic error message returned as a string   !// swagger:response errorResponse   "type errorResponseWrapper struct {   	// Description of the error   	// in: body   	Body GenericError   }       3// Validation errors defined as an array of strings   #// swagger:response errorValidation   $type errorValidationWrapper struct {   	// Collection of the errors   	// in: body   	Body ValidationError   }       // A list of products   $// swagger:response productsResponse   %type productsResponseWrapper struct {   	// All current products   	// in: body   	Body []data.Product   }       /// Data structure representing a single product   #// swagger:response productResponse   $type productResponseWrapper struct {   	// Newly created product   	// in: body   	Body data.Product   }       .// No content is returned by this API endpoint   %// swagger:response noContentResponse   &type noContentResponseWrapper struct {   }       1// swagger:parameters updateProduct createProduct   "type productParamsWrapper struct {   /	// Product data structure to Update or Create.   A	// Note: the id field is ignored by update and create operations   	// in: body   	// required: true   	Body data.Product   }       type ProductQueryParam struct {   :	// Currency used when returning the price of the product.   )	// when not specified is returned in GBP   	// in: query   	// required: false   	Currency string   }       5// swagger:parameters listSingleProduct deleteProduct   $type productIDParamsWrapper struct {   9	// The id of the product for which the operation relates   	// in: path   	// required: true   	ID int `json:"id"`   }5�_�                    E        ����                                                                                                                                                                                                                                                                                                                                                             _51     �   E   G   U       �   E   G   T    5�_�                     F   3    ����                                                                                                                                                                                                                                                                                                                                                             _5K    �               U   (// Package classification of Product API   //    // Documentation for Product API   //   //	Schemes: http   //	BasePath: /   //	Version: 1.0.0   //   //	Consumes:   //	- application/json   //   //	Produces:   //	- application/json   //   // swagger:meta   package handlers       Iimport "github.com/JamieBShaw/golang-mux-rest-api/products-rest-api/data"       //   A// NOTE: Types defined here are purely for documentation purposes   2// these types are not used by any of the handlers       -// Generic error message returned as a string   !// swagger:response errorResponse   "type errorResponseWrapper struct {   	// Description of the error   	// in: body   	Body GenericError   }       3// Validation errors defined as an array of strings   #// swagger:response errorValidation   $type errorValidationWrapper struct {   	// Collection of the errors   	// in: body   	Body ValidationError   }       // A list of products   $// swagger:response productsResponse   %type productsResponseWrapper struct {   	// All current products   	// in: body   	Body []data.Product   }       /// Data structure representing a single product   #// swagger:response productResponse   $type productResponseWrapper struct {   	// Newly created product   	// in: body   	Body data.Product   }       .// No content is returned by this API endpoint   %// swagger:response noContentResponse   &type noContentResponseWrapper struct {   }       1// swagger:parameters updateProduct createProduct   "type productParamsWrapper struct {   /	// Product data structure to Update or Create.   A	// Note: the id field is ignored by update and create operations   	// in: body   	// required: true   	Body data.Product   }       4// swagger:parameters listProducts listSingleProduct   type ProductQueryParam struct {   :	// Currency used when returning the price of the product.   )	// when not specified is returned in GBP   	// in: query   	// required: false   	Currency string   }       5// swagger:parameters listSingleProduct deleteProduct   $type productIDParamsWrapper struct {   9	// The id of the product for which the operation relates   	// in: path   	// required: true   	ID int `json:"id"`   }5��