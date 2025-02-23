// Code generated by cmd/blockhash; DO NOT EDIT.

package block

import "github.com/df-mc/dragonfly/server/world"

const (
	hashAir = iota
	hashAmethyst
	hashAncientDebris
	hashAndesite
	hashAnvil
	hashBanner
	hashBarrel
	hashBarrier
	hashBasalt
	hashBeacon
	hashBed
	hashBedrock
	hashBeetrootSeeds
	hashBlackstone
	hashBlastFurnace
	hashBlueIce
	hashBone
	hashBookshelf
	hashBrewingStand
	hashBricks
	hashCactus
	hashCake
	hashCalcite
	hashCampfire
	hashCarpet
	hashCarrot
	hashChain
	hashChest
	hashChiseledQuartz
	hashClay
	hashCoal
	hashCoalOre
	hashCobblestone
	hashCocoaBean
	hashComposter
	hashConcrete
	hashConcretePowder
	hashCopper
	hashCopperDoor
	hashCopperGrate
	hashCopperOre
	hashCopperTrapdoor
	hashCoral
	hashCoralBlock
	hashCraftingTable
	hashDeadBush
	hashDecoratedPot
	hashDeepslate
	hashDeepslateBricks
	hashDeepslateTiles
	hashDiamond
	hashDiamondOre
	hashDiorite
	hashDirt
	hashDirtPath
	hashDoubleFlower
	hashDoubleTallGrass
	hashDragonEgg
	hashDriedKelp
	hashDripstone
	hashEmerald
	hashEmeraldOre
	hashEnchantingTable
	hashEndBricks
	hashEndRod
	hashEndStone
	hashEnderChest
	hashFarmland
	hashFern
	hashFire
	hashFletchingTable
	hashFlower
	hashFroglight
	hashFurnace
	hashGlass
	hashGlassPane
	hashGlazedTerracotta
	hashGlowstone
	hashGold
	hashGoldOre
	hashGranite
	hashGrass
	hashGravel
	hashGrindstone
	hashHayBale
	hashHoneycomb
	hashHopper
	hashInvisibleBedrock
	hashIron
	hashIronBars
	hashIronOre
	hashItemFrame
	hashJukebox
	hashKelp
	hashLadder
	hashLantern
	hashLapis
	hashLapisOre
	hashLava
	hashLeaves
	hashLectern
	hashLight
	hashLilyPad
	hashLitPumpkin
	hashLog
	hashLoom
	hashMelon
	hashMelonSeeds
	hashMossCarpet
	hashMud
	hashMudBricks
	hashMuddyMangroveRoots
	hashNetherBrickFence
	hashNetherBricks
	hashNetherGoldOre
	hashNetherQuartzOre
	hashNetherSprouts
	hashNetherWart
	hashNetherWartBlock
	hashNetherite
	hashNetherrack
	hashNote
	hashObsidian
	hashPackedIce
	hashPackedMud
	hashPinkPetals
	hashPlanks
	hashPodzol
	hashPolishedBlackstoneBrick
	hashPolishedTuff
	hashPotato
	hashPrismarine
	hashPumpkin
	hashPumpkinSeeds
	hashPurpur
	hashPurpurPillar
	hashQuartz
	hashQuartzBricks
	hashQuartzPillar
	hashRawCopper
	hashRawGold
	hashRawIron
	hashReinforcedDeepslate
	hashResin
	hashResinBricks
	hashSand
	hashSandstone
	hashSeaLantern
	hashSeaPickle
	hashShortGrass
	hashShroomlight
	hashSign
	hashSkull
	hashSlab
	hashSmithingTable
	hashSmoker
	hashSnow
	hashSoulSand
	hashSoulSoil
	hashSponge
	hashSporeBlossom
	hashStainedGlass
	hashStainedGlassPane
	hashStainedTerracotta
	hashStairs
	hashStone
	hashStoneBricks
	hashStonecutter
	hashSugarCane
	hashTNT
	hashTerracotta
	hashTorch
	hashTuff
	hashTuffBricks
	hashVines
	hashWall
	hashWater
	hashWheatSeeds
	hashWood
	hashWoodDoor
	hashWoodFence
	hashWoodFenceGate
	hashWoodTrapdoor
	hashWool
	hashCustomBlockBase
)

// customBlockBase represents the base hash for all custom blocks.
var customBlockBase = uint64(hashCustomBlockBase - 1)

// NextHash returns the next free hash for custom blocks.
func NextHash() uint64 {
	customBlockBase++
	return customBlockBase
}

func (Air) Hash() (uint64, uint64) {
	return hashAir, 0
}

func (Amethyst) Hash() (uint64, uint64) {
	return hashAmethyst, 0
}

func (AncientDebris) Hash() (uint64, uint64) {
	return hashAncientDebris, 0
}

func (a Andesite) Hash() (uint64, uint64) {
	return hashAndesite, uint64(boolByte(a.Polished))
}

func (a Anvil) Hash() (uint64, uint64) {
	return hashAnvil, uint64(a.Type.Uint8()) | uint64(a.Facing)<<2
}

func (b Banner) Hash() (uint64, uint64) {
	return hashBanner, uint64(b.Attach.Uint8())
}

func (b Barrel) Hash() (uint64, uint64) {
	return hashBarrel, uint64(b.Facing) | uint64(boolByte(b.Open))<<3
}

func (Barrier) Hash() (uint64, uint64) {
	return hashBarrier, 0
}

func (b Basalt) Hash() (uint64, uint64) {
	return hashBasalt, uint64(boolByte(b.Polished)) | uint64(b.Axis)<<1
}

func (Beacon) Hash() (uint64, uint64) {
	return hashBeacon, 0
}

func (b Bed) Hash() (uint64, uint64) {
	return hashBed, uint64(b.Facing) | uint64(boolByte(b.HeadSide))<<2
}

func (b Bedrock) Hash() (uint64, uint64) {
	return hashBedrock, uint64(boolByte(b.InfiniteBurning))
}

func (b BeetrootSeeds) Hash() (uint64, uint64) {
	return hashBeetrootSeeds, uint64(b.Growth)
}

func (b Blackstone) Hash() (uint64, uint64) {
	return hashBlackstone, uint64(b.Type.Uint8())
}

func (b BlastFurnace) Hash() (uint64, uint64) {
	return hashBlastFurnace, uint64(b.Facing) | uint64(boolByte(b.Lit))<<2
}

func (BlueIce) Hash() (uint64, uint64) {
	return hashBlueIce, 0
}

func (b Bone) Hash() (uint64, uint64) {
	return hashBone, uint64(b.Axis)
}

func (Bookshelf) Hash() (uint64, uint64) {
	return hashBookshelf, 0
}

func (b BrewingStand) Hash() (uint64, uint64) {
	return hashBrewingStand, uint64(boolByte(b.LeftSlot)) | uint64(boolByte(b.MiddleSlot))<<1 | uint64(boolByte(b.RightSlot))<<2
}

func (Bricks) Hash() (uint64, uint64) {
	return hashBricks, 0
}

func (c Cactus) Hash() (uint64, uint64) {
	return hashCactus, uint64(c.Age)
}

func (c Cake) Hash() (uint64, uint64) {
	return hashCake, uint64(c.Bites)
}

func (Calcite) Hash() (uint64, uint64) {
	return hashCalcite, 0
}

func (c Campfire) Hash() (uint64, uint64) {
	return hashCampfire, uint64(c.Facing) | uint64(boolByte(c.Extinguished))<<2 | uint64(c.Type.Uint8())<<3
}

func (c Carpet) Hash() (uint64, uint64) {
	return hashCarpet, uint64(c.Colour.Uint8())
}

func (c Carrot) Hash() (uint64, uint64) {
	return hashCarrot, uint64(c.Growth)
}

func (c Chain) Hash() (uint64, uint64) {
	return hashChain, uint64(c.Axis)
}

func (c Chest) Hash() (uint64, uint64) {
	return hashChest, uint64(c.Facing)
}

func (ChiseledQuartz) Hash() (uint64, uint64) {
	return hashChiseledQuartz, 0
}

func (Clay) Hash() (uint64, uint64) {
	return hashClay, 0
}

func (Coal) Hash() (uint64, uint64) {
	return hashCoal, 0
}

func (c CoalOre) Hash() (uint64, uint64) {
	return hashCoalOre, uint64(c.Type.Uint8())
}

func (c Cobblestone) Hash() (uint64, uint64) {
	return hashCobblestone, uint64(boolByte(c.Mossy))
}

func (c CocoaBean) Hash() (uint64, uint64) {
	return hashCocoaBean, uint64(c.Facing) | uint64(c.Age)<<2
}

func (c Composter) Hash() (uint64, uint64) {
	return hashComposter, uint64(c.Level)
}

func (c Concrete) Hash() (uint64, uint64) {
	return hashConcrete, uint64(c.Colour.Uint8())
}

func (c ConcretePowder) Hash() (uint64, uint64) {
	return hashConcretePowder, uint64(c.Colour.Uint8())
}

func (c Copper) Hash() (uint64, uint64) {
	return hashCopper, uint64(c.Type.Uint8()) | uint64(c.Oxidation.Uint8())<<2 | uint64(boolByte(c.Waxed))<<4
}

func (d CopperDoor) Hash() (uint64, uint64) {
	return hashCopperDoor, uint64(d.Oxidation.Uint8()) | uint64(boolByte(d.Waxed))<<2 | uint64(d.Facing)<<3 | uint64(boolByte(d.Open))<<5 | uint64(boolByte(d.Top))<<6 | uint64(boolByte(d.Right))<<7
}

func (c CopperGrate) Hash() (uint64, uint64) {
	return hashCopperGrate, uint64(c.Oxidation.Uint8()) | uint64(boolByte(c.Waxed))<<2
}

func (c CopperOre) Hash() (uint64, uint64) {
	return hashCopperOre, uint64(c.Type.Uint8())
}

func (t CopperTrapdoor) Hash() (uint64, uint64) {
	return hashCopperTrapdoor, uint64(t.Oxidation.Uint8()) | uint64(boolByte(t.Waxed))<<2 | uint64(t.Facing)<<3 | uint64(boolByte(t.Open))<<5 | uint64(boolByte(t.Top))<<6
}

func (c Coral) Hash() (uint64, uint64) {
	return hashCoral, uint64(c.Type.Uint8()) | uint64(boolByte(c.Dead))<<3
}

func (c CoralBlock) Hash() (uint64, uint64) {
	return hashCoralBlock, uint64(c.Type.Uint8()) | uint64(boolByte(c.Dead))<<3
}

func (CraftingTable) Hash() (uint64, uint64) {
	return hashCraftingTable, 0
}

func (DeadBush) Hash() (uint64, uint64) {
	return hashDeadBush, 0
}

func (p DecoratedPot) Hash() (uint64, uint64) {
	return hashDecoratedPot, uint64(p.Facing)
}

func (d Deepslate) Hash() (uint64, uint64) {
	return hashDeepslate, uint64(d.Type.Uint8()) | uint64(d.Axis)<<2
}

func (d DeepslateBricks) Hash() (uint64, uint64) {
	return hashDeepslateBricks, uint64(boolByte(d.Cracked))
}

func (d DeepslateTiles) Hash() (uint64, uint64) {
	return hashDeepslateTiles, uint64(boolByte(d.Cracked))
}

func (Diamond) Hash() (uint64, uint64) {
	return hashDiamond, 0
}

func (d DiamondOre) Hash() (uint64, uint64) {
	return hashDiamondOre, uint64(d.Type.Uint8())
}

func (d Diorite) Hash() (uint64, uint64) {
	return hashDiorite, uint64(boolByte(d.Polished))
}

func (d Dirt) Hash() (uint64, uint64) {
	return hashDirt, uint64(boolByte(d.Coarse))
}

func (DirtPath) Hash() (uint64, uint64) {
	return hashDirtPath, 0
}

func (d DoubleFlower) Hash() (uint64, uint64) {
	return hashDoubleFlower, uint64(boolByte(d.UpperPart)) | uint64(d.Type.Uint8())<<1
}

func (d DoubleTallGrass) Hash() (uint64, uint64) {
	return hashDoubleTallGrass, uint64(boolByte(d.UpperPart)) | uint64(d.Type.Uint8())<<1
}

func (DragonEgg) Hash() (uint64, uint64) {
	return hashDragonEgg, 0
}

func (DriedKelp) Hash() (uint64, uint64) {
	return hashDriedKelp, 0
}

func (Dripstone) Hash() (uint64, uint64) {
	return hashDripstone, 0
}

func (Emerald) Hash() (uint64, uint64) {
	return hashEmerald, 0
}

func (e EmeraldOre) Hash() (uint64, uint64) {
	return hashEmeraldOre, uint64(e.Type.Uint8())
}

func (EnchantingTable) Hash() (uint64, uint64) {
	return hashEnchantingTable, 0
}

func (EndBricks) Hash() (uint64, uint64) {
	return hashEndBricks, 0
}

func (e EndRod) Hash() (uint64, uint64) {
	return hashEndRod, uint64(e.Facing)
}

func (EndStone) Hash() (uint64, uint64) {
	return hashEndStone, 0
}

func (c EnderChest) Hash() (uint64, uint64) {
	return hashEnderChest, uint64(c.Facing)
}

func (f Farmland) Hash() (uint64, uint64) {
	return hashFarmland, uint64(f.Hydration)
}

func (Fern) Hash() (uint64, uint64) {
	return hashFern, 0
}

func (f Fire) Hash() (uint64, uint64) {
	return hashFire, uint64(f.Type.Uint8()) | uint64(f.Age)<<1
}

func (FletchingTable) Hash() (uint64, uint64) {
	return hashFletchingTable, 0
}

func (f Flower) Hash() (uint64, uint64) {
	return hashFlower, uint64(f.Type.Uint8())
}

func (f Froglight) Hash() (uint64, uint64) {
	return hashFroglight, uint64(f.Type.Uint8()) | uint64(f.Axis)<<2
}

func (f Furnace) Hash() (uint64, uint64) {
	return hashFurnace, uint64(f.Facing) | uint64(boolByte(f.Lit))<<2
}

func (Glass) Hash() (uint64, uint64) {
	return hashGlass, 0
}

func (GlassPane) Hash() (uint64, uint64) {
	return hashGlassPane, 0
}

func (t GlazedTerracotta) Hash() (uint64, uint64) {
	return hashGlazedTerracotta, uint64(t.Colour.Uint8()) | uint64(t.Facing)<<4
}

func (Glowstone) Hash() (uint64, uint64) {
	return hashGlowstone, 0
}

func (Gold) Hash() (uint64, uint64) {
	return hashGold, 0
}

func (g GoldOre) Hash() (uint64, uint64) {
	return hashGoldOre, uint64(g.Type.Uint8())
}

func (g Granite) Hash() (uint64, uint64) {
	return hashGranite, uint64(boolByte(g.Polished))
}

func (Grass) Hash() (uint64, uint64) {
	return hashGrass, 0
}

func (Gravel) Hash() (uint64, uint64) {
	return hashGravel, 0
}

func (g Grindstone) Hash() (uint64, uint64) {
	return hashGrindstone, uint64(g.Attach.Uint8()) | uint64(g.Facing)<<2
}

func (h HayBale) Hash() (uint64, uint64) {
	return hashHayBale, uint64(h.Axis)
}

func (Honeycomb) Hash() (uint64, uint64) {
	return hashHoneycomb, 0
}

func (h Hopper) Hash() (uint64, uint64) {
	return hashHopper, uint64(h.Facing) | uint64(boolByte(h.Powered))<<3
}

func (InvisibleBedrock) Hash() (uint64, uint64) {
	return hashInvisibleBedrock, 0
}

func (Iron) Hash() (uint64, uint64) {
	return hashIron, 0
}

func (IronBars) Hash() (uint64, uint64) {
	return hashIronBars, 0
}

func (i IronOre) Hash() (uint64, uint64) {
	return hashIronOre, uint64(i.Type.Uint8())
}

func (i ItemFrame) Hash() (uint64, uint64) {
	return hashItemFrame, uint64(i.Facing) | uint64(boolByte(i.Glowing))<<3
}

func (Jukebox) Hash() (uint64, uint64) {
	return hashJukebox, 0
}

func (k Kelp) Hash() (uint64, uint64) {
	return hashKelp, uint64(k.Age)
}

func (l Ladder) Hash() (uint64, uint64) {
	return hashLadder, uint64(l.Facing)
}

func (l Lantern) Hash() (uint64, uint64) {
	return hashLantern, uint64(boolByte(l.Hanging)) | uint64(l.Type.Uint8())<<1
}

func (Lapis) Hash() (uint64, uint64) {
	return hashLapis, 0
}

func (l LapisOre) Hash() (uint64, uint64) {
	return hashLapisOre, uint64(l.Type.Uint8())
}

func (l Lava) Hash() (uint64, uint64) {
	return hashLava, uint64(boolByte(l.Still)) | uint64(l.Depth)<<1 | uint64(boolByte(l.Falling))<<9
}

func (l Leaves) Hash() (uint64, uint64) {
	return hashLeaves, uint64(l.Wood.Uint8()) | uint64(boolByte(l.Persistent))<<4 | uint64(boolByte(l.ShouldUpdate))<<5
}

func (l Lectern) Hash() (uint64, uint64) {
	return hashLectern, uint64(l.Facing)
}

func (l Light) Hash() (uint64, uint64) {
	return hashLight, uint64(l.Level)
}

func (LilyPad) Hash() (uint64, uint64) {
	return hashLilyPad, 0
}

func (l LitPumpkin) Hash() (uint64, uint64) {
	return hashLitPumpkin, uint64(l.Facing)
}

func (l Log) Hash() (uint64, uint64) {
	return hashLog, uint64(l.Wood.Uint8()) | uint64(boolByte(l.Stripped))<<4 | uint64(l.Axis)<<5
}

func (l Loom) Hash() (uint64, uint64) {
	return hashLoom, uint64(l.Facing)
}

func (Melon) Hash() (uint64, uint64) {
	return hashMelon, 0
}

func (m MelonSeeds) Hash() (uint64, uint64) {
	return hashMelonSeeds, uint64(m.Growth) | uint64(m.Direction)<<8
}

func (MossCarpet) Hash() (uint64, uint64) {
	return hashMossCarpet, 0
}

func (Mud) Hash() (uint64, uint64) {
	return hashMud, 0
}

func (MudBricks) Hash() (uint64, uint64) {
	return hashMudBricks, 0
}

func (m MuddyMangroveRoots) Hash() (uint64, uint64) {
	return hashMuddyMangroveRoots, uint64(m.Axis)
}

func (NetherBrickFence) Hash() (uint64, uint64) {
	return hashNetherBrickFence, 0
}

func (n NetherBricks) Hash() (uint64, uint64) {
	return hashNetherBricks, uint64(n.Type.Uint8())
}

func (NetherGoldOre) Hash() (uint64, uint64) {
	return hashNetherGoldOre, 0
}

func (NetherQuartzOre) Hash() (uint64, uint64) {
	return hashNetherQuartzOre, 0
}

func (NetherSprouts) Hash() (uint64, uint64) {
	return hashNetherSprouts, 0
}

func (n NetherWart) Hash() (uint64, uint64) {
	return hashNetherWart, uint64(n.Age)
}

func (n NetherWartBlock) Hash() (uint64, uint64) {
	return hashNetherWartBlock, uint64(boolByte(n.Warped))
}

func (Netherite) Hash() (uint64, uint64) {
	return hashNetherite, 0
}

func (Netherrack) Hash() (uint64, uint64) {
	return hashNetherrack, 0
}

func (Note) Hash() (uint64, uint64) {
	return hashNote, 0
}

func (o Obsidian) Hash() (uint64, uint64) {
	return hashObsidian, uint64(boolByte(o.Crying))
}

func (PackedIce) Hash() (uint64, uint64) {
	return hashPackedIce, 0
}

func (PackedMud) Hash() (uint64, uint64) {
	return hashPackedMud, 0
}

func (p PinkPetals) Hash() (uint64, uint64) {
	return hashPinkPetals, uint64(p.AdditionalCount) | uint64(p.Facing)<<8
}

func (p Planks) Hash() (uint64, uint64) {
	return hashPlanks, uint64(p.Wood.Uint8())
}

func (Podzol) Hash() (uint64, uint64) {
	return hashPodzol, 0
}

func (b PolishedBlackstoneBrick) Hash() (uint64, uint64) {
	return hashPolishedBlackstoneBrick, uint64(boolByte(b.Cracked))
}

func (PolishedTuff) Hash() (uint64, uint64) {
	return hashPolishedTuff, 0
}

func (p Potato) Hash() (uint64, uint64) {
	return hashPotato, uint64(p.Growth)
}

func (p Prismarine) Hash() (uint64, uint64) {
	return hashPrismarine, uint64(p.Type.Uint8())
}

func (p Pumpkin) Hash() (uint64, uint64) {
	return hashPumpkin, uint64(boolByte(p.Carved)) | uint64(p.Facing)<<1
}

func (p PumpkinSeeds) Hash() (uint64, uint64) {
	return hashPumpkinSeeds, uint64(p.Growth) | uint64(p.Direction)<<8
}

func (Purpur) Hash() (uint64, uint64) {
	return hashPurpur, 0
}

func (p PurpurPillar) Hash() (uint64, uint64) {
	return hashPurpurPillar, uint64(p.Axis)
}

func (q Quartz) Hash() (uint64, uint64) {
	return hashQuartz, uint64(boolByte(q.Smooth))
}

func (QuartzBricks) Hash() (uint64, uint64) {
	return hashQuartzBricks, 0
}

func (q QuartzPillar) Hash() (uint64, uint64) {
	return hashQuartzPillar, uint64(q.Axis)
}

func (RawCopper) Hash() (uint64, uint64) {
	return hashRawCopper, 0
}

func (RawGold) Hash() (uint64, uint64) {
	return hashRawGold, 0
}

func (RawIron) Hash() (uint64, uint64) {
	return hashRawIron, 0
}

func (ReinforcedDeepslate) Hash() (uint64, uint64) {
	return hashReinforcedDeepslate, 0
}

func (Resin) Hash() (uint64, uint64) {
	return hashResin, 0
}

func (r ResinBricks) Hash() (uint64, uint64) {
	return hashResinBricks, uint64(boolByte(r.Chiseled))
}

func (s Sand) Hash() (uint64, uint64) {
	return hashSand, uint64(boolByte(s.Red))
}

func (s Sandstone) Hash() (uint64, uint64) {
	return hashSandstone, uint64(s.Type.Uint8()) | uint64(boolByte(s.Red))<<2
}

func (SeaLantern) Hash() (uint64, uint64) {
	return hashSeaLantern, 0
}

func (s SeaPickle) Hash() (uint64, uint64) {
	return hashSeaPickle, uint64(s.AdditionalCount) | uint64(boolByte(s.Dead))<<8
}

func (ShortGrass) Hash() (uint64, uint64) {
	return hashShortGrass, 0
}

func (Shroomlight) Hash() (uint64, uint64) {
	return hashShroomlight, 0
}

func (s Sign) Hash() (uint64, uint64) {
	return hashSign, uint64(s.Wood.Uint8()) | uint64(s.Attach.Uint8())<<4
}

func (s Skull) Hash() (uint64, uint64) {
	return hashSkull, uint64(s.Type.Uint8()) | uint64(s.Attach.FaceUint8())<<3
}

func (s Slab) Hash() (uint64, uint64) {
	return hashSlab, world.BlockHash(s.Block) | uint64(boolByte(s.Top))<<32 | uint64(boolByte(s.Double))<<33
}

func (SmithingTable) Hash() (uint64, uint64) {
	return hashSmithingTable, 0
}

func (s Smoker) Hash() (uint64, uint64) {
	return hashSmoker, uint64(s.Facing) | uint64(boolByte(s.Lit))<<2
}

func (Snow) Hash() (uint64, uint64) {
	return hashSnow, 0
}

func (SoulSand) Hash() (uint64, uint64) {
	return hashSoulSand, 0
}

func (SoulSoil) Hash() (uint64, uint64) {
	return hashSoulSoil, 0
}

func (s Sponge) Hash() (uint64, uint64) {
	return hashSponge, uint64(boolByte(s.Wet))
}

func (SporeBlossom) Hash() (uint64, uint64) {
	return hashSporeBlossom, 0
}

func (g StainedGlass) Hash() (uint64, uint64) {
	return hashStainedGlass, uint64(g.Colour.Uint8())
}

func (p StainedGlassPane) Hash() (uint64, uint64) {
	return hashStainedGlassPane, uint64(p.Colour.Uint8())
}

func (t StainedTerracotta) Hash() (uint64, uint64) {
	return hashStainedTerracotta, uint64(t.Colour.Uint8())
}

func (s Stairs) Hash() (uint64, uint64) {
	return hashStairs, world.BlockHash(s.Block) | uint64(boolByte(s.UpsideDown))<<32 | uint64(s.Facing)<<33
}

func (s Stone) Hash() (uint64, uint64) {
	return hashStone, uint64(boolByte(s.Smooth))
}

func (s StoneBricks) Hash() (uint64, uint64) {
	return hashStoneBricks, uint64(s.Type.Uint8())
}

func (s Stonecutter) Hash() (uint64, uint64) {
	return hashStonecutter, uint64(s.Facing)
}

func (c SugarCane) Hash() (uint64, uint64) {
	return hashSugarCane, uint64(c.Age)
}

func (TNT) Hash() (uint64, uint64) {
	return hashTNT, 0
}

func (Terracotta) Hash() (uint64, uint64) {
	return hashTerracotta, 0
}

func (t Torch) Hash() (uint64, uint64) {
	return hashTorch, uint64(t.Facing) | uint64(t.Type.Uint8())<<3
}

func (t Tuff) Hash() (uint64, uint64) {
	return hashTuff, uint64(boolByte(t.Chiseled))
}

func (t TuffBricks) Hash() (uint64, uint64) {
	return hashTuffBricks, uint64(boolByte(t.Chiseled))
}

func (v Vines) Hash() (uint64, uint64) {
	return hashVines, uint64(boolByte(v.NorthDirection)) | uint64(boolByte(v.EastDirection))<<1 | uint64(boolByte(v.SouthDirection))<<2 | uint64(boolByte(v.WestDirection))<<3
}

func (w Wall) Hash() (uint64, uint64) {
	return hashWall, world.BlockHash(w.Block) | uint64(w.NorthConnection.Uint8())<<32 | uint64(w.EastConnection.Uint8())<<34 | uint64(w.SouthConnection.Uint8())<<36 | uint64(w.WestConnection.Uint8())<<38 | uint64(boolByte(w.Post))<<40
}

func (w Water) Hash() (uint64, uint64) {
	return hashWater, uint64(boolByte(w.Still)) | uint64(w.Depth)<<1 | uint64(boolByte(w.Falling))<<9
}

func (s WheatSeeds) Hash() (uint64, uint64) {
	return hashWheatSeeds, uint64(s.Growth)
}

func (w Wood) Hash() (uint64, uint64) {
	return hashWood, uint64(w.Wood.Uint8()) | uint64(boolByte(w.Stripped))<<4 | uint64(w.Axis)<<5
}

func (d WoodDoor) Hash() (uint64, uint64) {
	return hashWoodDoor, uint64(d.Wood.Uint8()) | uint64(d.Facing)<<4 | uint64(boolByte(d.Open))<<6 | uint64(boolByte(d.Top))<<7 | uint64(boolByte(d.Right))<<8
}

func (w WoodFence) Hash() (uint64, uint64) {
	return hashWoodFence, uint64(w.Wood.Uint8())
}

func (f WoodFenceGate) Hash() (uint64, uint64) {
	return hashWoodFenceGate, uint64(f.Wood.Uint8()) | uint64(f.Facing)<<4 | uint64(boolByte(f.Open))<<6 | uint64(boolByte(f.Lowered))<<7
}

func (t WoodTrapdoor) Hash() (uint64, uint64) {
	return hashWoodTrapdoor, uint64(t.Wood.Uint8()) | uint64(t.Facing)<<4 | uint64(boolByte(t.Open))<<6 | uint64(boolByte(t.Top))<<7
}

func (w Wool) Hash() (uint64, uint64) {
	return hashWool, uint64(w.Colour.Uint8())
}
