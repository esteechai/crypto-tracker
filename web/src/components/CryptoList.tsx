import React from 'react'
import { coinbaseProducts, coinbaseTicker } from '../struct';
import { StoreContainer } from '../store';
import { Card, Icon, Modal } from 'semantic-ui-react';
import {Link, Redirect} from 'react-router-dom'

interface coinbaseProductsProps {
    product: coinbaseProducts
}

const CryptoList: React.FC<coinbaseProductsProps> = () => {
    const store = StoreContainer.useContainer()
       return (
        <Card.Group>
            {store.searchResult?store.searchResult.map((product:coinbaseProducts)=>{
               return ( 
                    <Card key={product.id}>
                    <Card.Header>{product.id}</Card.Header>             
                    <Card.Meta>Base currency: {product.base_currency}</Card.Meta>
                    <Card.Meta>Quote currency: {product.quote_currency}</Card.Meta>
                    <Card.Content extra>
                            <span className="right floated">
                                <Modal trigger={<button className="tiny ui button" onClick={() => {store.handleSelectedProduct(product.id)}}>More details
                                        </button>} size="tiny">
                                            {store.ticker?<React.Fragment><Modal.Header>{store.ticker.ID}</Modal.Header>
                                        <Modal.Description>
                                        <p>Price: {Number(store.ticker.Price).toFixed(2)}</p>
                                        <p>Ask: {Number(store.ticker.Ask).toFixed(2)}</p>
                                        </Modal.Description></React.Fragment>:null}
                                </Modal>                                   
                            </span>
                        <span onClick={() => {store.handleFavourite(product.id, store.currentUser)}}>
                            {store.handleFavIcon(product.id) ? <i className="large red heart icon"></i> : <i className="large red heart outline icon" ></i> }
                        </span>     
                    </Card.Content>    
                </Card>            
               )             
           }):null}
        </Card.Group>                        
    )
}

export default CryptoList
