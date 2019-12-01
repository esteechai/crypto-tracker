import React from 'react'
import { coinbaseProducts } from '../struct';
import { StoreContainer } from '../store';
import { Card, Icon } from 'semantic-ui-react';

interface coinbaseProductsProps {
    product: coinbaseProducts
}

const CryptoList: React.FC<coinbaseProductsProps> = () => {
    const store = StoreContainer.useContainer()

    return (
    <Card.Group>
            {store.productList?store.productList.map((product:coinbaseProducts)=>{
               return ( 
                <Card>
               <Card.Header>{product.id}</Card.Header>             
               <Card.Meta>Base currency: {product.base_currency}</Card.Meta>
               <Card.Meta>Quote currency: {product.quote_currency}</Card.Meta>
               <Card.Content extra>
                    <span className="right floated"><button className="tiny ui button">More details</button></span>
                   <span><Icon name="heart outline" color="red" link onClick=""/></span> 
               </Card.Content>    
               </Card>                                                  
               )             
           }):null}
        </Card.Group>                        
    )
}

export default CryptoList