import React from 'react'
import { StoreContainer } from '../store'
import { UserFavList } from '../struct'
import { Card} from 'semantic-ui-react';
import { Redirect } from 'react-router'

interface FavCryptoProps {
    favProducts: UserFavList
}

const FavCrypto: React.FC<FavCryptoProps> = () => {
    const store = StoreContainer.useContainer() 
    if (!store.isLogin){
        return <Redirect to= '/login' />
    }
    return (
        <Card.Group>  
        {store.userFavList?store.userFavList.map((favProducts:UserFavList)=>{           
            return (
                <Card key={favProducts.ID}>
                    <Card.Header>{favProducts.ID}</Card.Header>  
                    <Card.Meta>Last update: {favProducts.Time}</Card.Meta>           
                    <Card.Description>Bid: {favProducts.Bid}</Card.Description>
                    <Card.Description>Ask: {favProducts.Ask}</Card.Description>
                    <Card.Description>Price: {favProducts.Price}</Card.Description>
                    <Card.Description>Size: {Number(favProducts.Size).toFixed(2)}</Card.Description>
                    <Card.Description>Volume: {Number(favProducts.Volume).toFixed(2)}</Card.Description>
                    <Card.Content extra>
                        <span onClick={() => {store.handleFavourite(favProducts.ID, store.currentUser)}}>
                            <i className="large red heart link icon"></i> 
                        </span>     
                    </Card.Content>    
                </Card>  
           )            
        }):null} 
        </Card.Group>   
    )
}

export {FavCrypto}
