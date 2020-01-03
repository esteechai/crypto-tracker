import React from 'react'
import { coinbaseProducts} from '../struct';
import { StoreContainer } from '../store';
import { Card, Modal } from 'semantic-ui-react';
import {Redirect} from 'react-router-dom'

interface coinbaseProductsProps {
    product: coinbaseProducts
}

const CryptoList: React.FC<coinbaseProductsProps> = () => {
    const store = StoreContainer.useContainer()
    if (!store.isLogin){
        return <Redirect to= '/login' />
    }
       return (
            <Card.Group className="ui centered">
                {store.searchResult?store.searchResult.map((product:coinbaseProducts)=>{
                return ( 
                        <Card key={product.ID}>
                        <Card.Header>{product.ID}</Card.Header>             
                        <Card.Meta>Base currency: {product.BaseCurrency}</Card.Meta>
                        <Card.Meta>Quote currency: {product.QuoteCurrency}</Card.Meta>
                        <Card.Content extra>
                                <span className="right floated">
                                    <Modal trigger={<button className="tiny ui button" onClick={() => {store.handleSelectedProduct(product.ID)}}>More details
                                            </button>} size="tiny">
                                                {store.ticker?<React.Fragment><Modal.Header>{store.ticker.ID}</Modal.Header>
                                            <Modal.Description>
                                            <p>Price: {Number(store.ticker.Price).toFixed(2)}</p>
                                            <p>Ask: {Number(store.ticker.Ask).toFixed(2)}</p>
                                            </Modal.Description></React.Fragment>:null}
                                    </Modal>                                   
                                </span>
                            <span onClick={() => {store.handleFavourite(product.ID, store.currentUser)}}>
                                {store.handleFavIcon(product.ID) ? <i className="large red heart link icon"></i> : <i className="large red heart outline link icon" ></i> }
                            </span>     
                        </Card.Content>    
                    </Card>            
                )             
            }):null}
            </Card.Group>   
    )
}

export {CryptoList}
