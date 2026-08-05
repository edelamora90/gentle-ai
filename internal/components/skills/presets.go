package skills

import "github.com/gentleman-programming/gentle-ai/v2/internal/model"

// sddSkills are the SDD orchestrator skills — always included.
var sddSkills = []model.SkillID{
	model.SkillSDDInit,
	model.SkillSDDExplore,
	model.SkillSDDPropose,
	model.SkillSDDSpec,
	model.SkillSDDDesign,
	model.SkillSDDVisual,
	model.SkillSDDTasks,
	model.SkillSDDApply,
	model.SkillSDDVerify,
	model.SkillSDDArchive,
	model.SkillSDDOnboard,
	model.SkillJudgmentDay,
}

// designSkills back the sdd-visual phase. They ship with the recommended and
// full tiers, not with "minimal" — that preset is SDD-only by contract. In a
// minimal install sdd-visual still runs, but it works from its own phase
// guidance instead of delegating to these sub-skills.
var designSkills = []model.SkillID{
	model.SkillDesignResearcher,
	model.SkillFrontendDesign,
	model.SkillCopywritingForUI,
	model.SkillImageSourcingPolicy,
	model.SkillAccessibilityBaseline,
	model.SkillVisualCritic,
}

// foundationSkills are baseline learning skills for the "recommended" tier.
var foundationSkills = []model.SkillID{
	model.SkillGoTesting,
	model.SkillGentleAIBench,
	model.SkillCreator,
	model.SkillImprover,
	model.SkillBranchPR,
	model.SkillIssueCreation,
	model.SkillSkillRegistry,
	model.SkillChainedPR,
	model.SkillCognitiveDoc,
	model.SkillCommentWriter,
	model.SkillWorkUnitCommits,
	model.SkillRDDDefectWorkflow,
	model.SkillSystemicIssueTriage,
}

// SkillsForPreset returns which skills should be installed for a given preset.
//
//   - "minimal" / PresetMinimal:       SDD skills only
//   - "ecosystem-only" / PresetEcosystemOnly: SDD + common framework skills
//   - "full-gentleman" / PresetFullGentleman: all available skills
//   - "custom" / PresetCustom:         empty (caller should provide explicit list)
func SkillsForPreset(preset model.PresetID) []model.SkillID {
	switch preset {
	case model.PresetMinimal:
		return concatSkills(sddSkills)
	case model.PresetEcosystemOnly:
		return concatSkills(sddSkills, designSkills, foundationSkills)
	case model.PresetFullGentleman:
		return concatSkills(sddSkills, designSkills, foundationSkills)
	case model.PresetCustom:
		return nil
	default:
		// Unknown preset — default to full.
		return concatSkills(sddSkills, designSkills, foundationSkills)
	}
}

// AllSkillIDs returns every known skill ID.
func AllSkillIDs() []model.SkillID {
	return concatSkills(sddSkills, designSkills, foundationSkills)
}

// concatSkills returns a fresh slice holding every group in order. It never
// appends into a package-level slice, so callers cannot alias shared backing
// arrays.
func concatSkills(groups ...[]model.SkillID) []model.SkillID {
	total := 0
	for _, group := range groups {
		total += len(group)
	}
	all := make([]model.SkillID, 0, total)
	for _, group := range groups {
		all = append(all, group...)
	}
	return all
}
