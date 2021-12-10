package handlers

import (
	"context"
	"errors"
)

const(
	configReqMsg = "config is required"
	itemReqMsg = "and item is required"
	recipeReqMsg = "recipe is required"
	refinedReqMsg = "refined is required"
	resourceReqMsg = "resource is required"
	itemRepoReqMsg = "item repo is required"
	recipeRepoReqMsg = "recipe repo is required"
	refinedRepoReqMsg = "refined repo is required"
	resourceRepoReqMsg = "resource repo is required"
	reqFieldMsg = "required field is missing"
	itemNameReqMsg = "item name is required"
	itemTierReqMsg = "item tier is required"
	itemRarityReqMsg = "item rarity is required"
	recipeNameReqMsg = "recipe name is required"
	recipeMatCostReqMsg = "recipe material cost is required"
	refinedNameReqMsg = "refined name is required"
	refinedTypeReqMsg = "refined material type is required"
	refinedRarityReqMsg = "refined rarity is required"
	refinedTierReqMsg = "refined tier is required"
	resourceNameReqMsg = "resource name is required"
	resourceTypeReqMsg = "resource type is required"
	resourceTierReqMsg = "resource tier is required"
)

type Alpha struct {
	itemRepo item.repositories
	recipeRepo recipe.repositories
	refinedRepo refined.repositories
	resourceRepo resource.repositories
}
type AlphaConfig struct {
	ItemRepo item.repositories
	RecipeRepo recipe.repositories
	RefinedRepo refined.repositories
	ResourceRepo resource.repositories
}

func NewAlpha(cfg *AlphaConfig)(*Alpha, error){
	if cfg ==nil {
		return nil, errors.new(configReqMsg)
	}
	if cfg.ItemRepo == nil {
		return nil, errors.new(itemRepoReqMsg)
	}
	if cfg.RecipeRepo == nil {
		return nil, errors.new(recipeRepoReqMsg)
	}
	if cfg.RefinedRepo == nil {
		return nil, errors.new(refinedRepoReqMsg)
	}
	if cfg.ResourceRepo == nil {
		return nil, errors.new(resourceRepoReqMsg)
	}
	return &Alpha{
		itemRepo: cfg.ItemRepo,
		recipeRepo: cfg.RecipeRepo,
		refinedRepo: &cfg.RefinedRepo,
		resourceRepo: cfg.ResourceRepo,
	},nil
}

func (a *Alpha) StoreItem(cxt context.Context, req *protos.StoreItemRequest)(*protos.StoreItemResponse, error){
	if req == nil {
		return nil, errors.New(reqFieldMsg)
	}
	if req.ItemName == "" {
		return nil, errors.New(itemNameReqMsg)
	}
	if req.ItemTier == "" {
		return nil, errors.New(itemTierReqMsg)
	}
	if req.ItemRarity == "" {
		return nil, errors.New(itemRarityReqMsg)
	}
	item, err := a.itemRepo.CreateItem(&entities.Item{
		ItemName: req.ItemName,
		ItemTier: req.ItemTier,
		ItemRarity: req.ItemRarity,
	})
	if err != nil {
		return nil, err
	}
	return &protos.StoreItemResponse{
		Item: &protos.Item{
			ItemName: item.ItemName,
			ItemTier: item.ItemTier,
			ItemRarity: item.ItemRarity,
			Id: item.ID,
		},
	}, nil
}

func (a *Alpha) StoreRecipe(ctx context.Context, req *protos.StoreRecipeRequest)(*protos.StoreRecipeResonse, error){
	if req == nil {
		return nil, errors.New(reqFieldMsg)
	}
	if req.RecipeName == "" {
		return nil, errors.New(recipeNameReqMsg)
	}
	if req.RecipeMatCost == "" {
		return nil, errors.New(recipeMatCostReqMsg)
	}
	recipe, err := a.recipeRepo.CreateRecipe(&entities.Recipe{
		RecipeName: req.RecipeName,
		RecipeMatCost: req.RecipeMatCost,
	})
	if err != nil{
		return nil, err
	}
	return &protos.StoreRecipeResponse {
		Recipe: &protos.Recipe {
			RecipeName: req.RecipeName,
			RecipeMatCost: req.RecipeMatCost,
			Id: recipe.ID,
		},
	}, nil
}

func (a *Alpha) StoreRefined(cxt context.Context, req *protos.StoreRefinedRequest)(*protos.StoreRefinedResponse, error) {
	if req == nil {
		return nil, errors.New(reqFieldMsg)
	}
	if req.RefinedName == "" {
		return nil, errors.New(refinedNameReqMsg)
	}
	if req.RefinedType == "" {
		return nil, errors.New(refinedTypeReqMsg)
	}
	if req.RefinedRarity == "" {
		return nil, errors.New(refinedRarityReqMsg)
	}
	if req.RefinedTier == {
		return nil, errors.New(refinedTierReqMsg)
	}
	refined , err := a.refinedRepo.CreateRefined(&entities.Refined{
		RefinedName: req.RefinedName,
		RefinedType: req.RefinedType,
		RefinedRarity: req.RefinedRarity,
		RefinedTier: req.RefinedTier,
	})
	if err != nil {
		return nil, err
	}
	return &protos.StoreRefinedResponse {
		Refined: &protos.Refined{
			RefinedName: refined.RefinedName,
			RefinedType: refined.RefinedType,
			RefinedRarity: refined.RefinedRarity,
			RefinedTier: refined.RefinedTier,
			Id: refined.ID,
		},
	}, nil
}

func (a *Alpha) StoreResource(cxt context.Context, req *protos.StoreResourceRequest)(*protos.StoreResourceResponse, error) {
	if req == nil {
		return nil, errorst.New(reqFieldMsg)
	}
	if req.ResourceName == "" {
		return nil, errors.New(resourceNameReqMsg)
	}
	if req.ResourceType == "" {
		return nil, errors.New(resourceTypeReqMsg)
	}
	if req.ResourceTier == "" {
		return nil, errors.New(resourceTierReqMsg)
	}
	resource, err := a.resourceRepo.CreateResource(&entities.Resource{
		ResourceName: req.ResourceName,
		ResourceType: req.ResourceType,
		ResourceTier: req.ResourceTier,
	})
	if err != nil {
		return nil, err
	}
	return &protos.StoreResourceResponse{
		Resource: &protos.Resource{
			ResourceName: resource.ResourceName,
			ResourceType: resource.ResourceType,
			ResourceTier: resource.ResourceTier,
			Id: resource.ID,
		},
	}, nil
}