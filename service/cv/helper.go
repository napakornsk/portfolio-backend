package cv

import (
	pb "github.com/napakornsk/cv-backend/pb"
	"go.mongodb.org/mongo-driver/bson"
)

// Helper function to build WorkExperience array
func buildWorkExperiences(workExps []*pb.WorkExperience) []bson.D {
    var bsonWorkExps []bson.D
    for _, exp := range workExps {
        bsonWorkExps = append(bsonWorkExps, bson.D{
            {"company_name", exp.GetCompanyName()},
            {"role", exp.GetRole()},
            {"description", exp.GetDescription()},
            {"start_date", exp.GetStartDate()},
            {"end_date", exp.GetEndDate()},
            {"work_desc", buildWorkDescriptions(exp.GetWorkDesc())},
        })
    }
    return bsonWorkExps
}

// Helper function to build WorkDesc array
func buildWorkDescriptions(workDescs []*pb.WorkDesc) []bson.D {
    var bsonWorkDescs []bson.D
    for _, desc := range workDescs {
        bsonWorkDescs = append(bsonWorkDescs, bson.D{
            {"description", desc.GetDescription()},
        })
    }
    return bsonWorkDescs
}

// Helper function to build Education array
func buildEducationEntries(educationEntries []*pb.Education) []bson.D {
    var bsonEdus []bson.D
    for _, edu := range educationEntries {
        bsonEdus = append(bsonEdus, bson.D{
            {"school_name", edu.GetSchoolName()},
            {"degree", edu.GetDegree()},
            {"field", edu.GetField()},
            {"start_date", edu.GetStartDate()},
            {"end_date", edu.GetEndDate()},
        })
    }
    return bsonEdus
}

// Helper function to build Skills array
func buildSkills(skills []*pb.Skill) []bson.D {
    var bsonSkills []bson.D
    for _, skill := range skills {
        bsonSkills = append(bsonSkills, bson.D{
            {"skill_name", skill.GetSkillName()},
        })
    }
    return bsonSkills
}